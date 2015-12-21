package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/drone-plugins/drone-git-push/repo"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build     string
	buildDate string
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
	os.Setenv("GIT_SSH_KEY", sshPrivateKeyPath())

	tasks := strings.Fields(vargs.Tasks)

	if len(tasks) == 0 {
		fmt.Println("Please provide Capistrano tasks to execute")
		os.Exit(1)
		return
	}

	log("Running Bundler")
	bundle_args := []string{"install", "--quiet"}
	if len(vargs.BundlePath) > 0 {
		bundle_args = append(bundle_args, "--path", vargs.BundlePath)
	}
	bundle := dw.bundle(bundle_args...)
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

func sshPrivateKeyPath() string {
	home := "/root"
	if currentUser, err := user.Current(); err == nil {
		home = currentUser.HomeDir
	}
	return filepath.Join(home, ".ssh", "id_rsa")
}

func log(message string, a ...interface{}) {
	fmt.Printf("=> %s\n", fmt.Sprintf(message, a...))
}
