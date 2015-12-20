package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build             string
	buildDate         string
	sshKeyPath        string = "/root/.ssh"
	sshPrivateKeyPath string = path.Join(sshKeyPath, "id_rsa")
	sshPublicKeyPath  string = path.Join(sshKeyPath, "id_rsa.pub")
)

type DeployWorkspace struct {
  Workspace drone.Workspace
}

func main() {
	fmt.Printf("Drone Capistrano Plugin built at %s\n", buildDate)

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

  dw := DeployWorkspace{workspace}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	fmt.Printf("Installing your deploy key to %s\n", sshKeyPath)
	os.MkdirAll(sshKeyPath, 0700)
	ioutil.WriteFile(sshPrivateKeyPath, []byte(workspace.Keys.Private), 0600)
	ioutil.WriteFile(sshPublicKeyPath, []byte(workspace.Keys.Public), 0644)

  os.Setenv("BUILD_PATH", workspace.Path)
	os.Setenv("GIT_SSH_KEY", sshPrivateKeyPath)

	tasks := strings.Fields(vargs.Tasks)

	if len(tasks) == 0 {
		fmt.Println("Please provide Capistrano tasks to execute")
		os.Exit(1)
		return
	}

  fmt.Printf("Running Bundler\n")
  bundle := dw.bundle("install")
  if err := bundle.Run(); err != nil {
    fmt.Println(err)
    os.Exit(1)
    return
  }

  fmt.Printf("Running Capistrano\n")
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
