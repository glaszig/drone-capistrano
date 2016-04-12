package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/drone-plugins/drone-git-push/repo"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build          string
	buildDate      string
	privateKeyPath string = "/root/.ssh/id_rsa"
)

type DeployWorkspace struct {
	Workspace drone.Workspace
}

func main() {
	fmt.Printf("Drone Capistrano Plugin built at %s\n", buildDate)

	workspace := drone.Workspace{}
	vargs := Params{}

	dw := DeployWorkspace{workspace}

	plugin.Param("workspace", &workspace)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	log("Installing Drone's ssh key")
	if err := repo.WriteKey(&workspace); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Setenv("BUILD_PATH", workspace.Path)
	os.Setenv("GIT_SSH_KEY", privateKeyPath)

	if vargs.Debug {
		bundleAppConfig := os.Getenv("BUNDLE_APP_CONFIG")
		fmt.Printf("BUNDLE_APP_CONFIG: %s\n", bundleAppConfig)
	}

	tasks := strings.Fields(vargs.Tasks)

	if len(tasks) == 0 {
		fmt.Println("Please provide Capistrano tasks to execute")
		os.Exit(1)
		return
	}

	log("Running Bundler")
	bundle := dw.bundle(bundlerArgs(vargs)...)
	if err := bundle.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	log("Running Capistrano")
	capistrano := dw.cap(tasks...)
	if err := capistrano.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func bundlerArgs(vargs Params) []string {
	args := []string{"install"}
	if ! vargs.Debug {
		args = append(args, "--quiet")
	}
	if len(vargs.BundlePath) > 0 {
		args = append(args, "--path", vargs.BundlePath)
	}
	return args
}

func (w *DeployWorkspace) cap(tasks ...string) *exec.Cmd {
	args := append([]string{"exec", "cap"}, tasks...)
	return w.bundle(args...)
}

func (w *DeployWorkspace) bundle(args ...string) *exec.Cmd {
	return w.command("/bundle.sh", args...)
}

func (w *DeployWorkspace) command(cmd string, args ...string) *exec.Cmd {
	c := exec.Command(cmd, args...)
	c.Dir = w.Workspace.Path
	c.Env = os.Environ()
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}

func log(message string, a ...interface{}) {
	fmt.Printf("=> %s\n", fmt.Sprintf(message, a...))
}
