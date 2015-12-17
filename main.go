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

func main() {
	fmt.Printf("Drone Capistrano Plugin built at %s\n", buildDate)

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	fmt.Printf("Installing your deploy key to %s\n", sshKeyPath)
	os.MkdirAll(sshKeyPath, 0700)
	ioutil.WriteFile(sshPrivateKeyPath, []byte(workspace.Keys.Private), 0600)
	ioutil.WriteFile(sshPublicKeyPath, []byte(workspace.Keys.Public), 0644)

	// set private key to use with $GIT_SSH wrapper
	os.Setenv("GIT_SSH", "/git_ssh.sh")
	os.Setenv("GIT_SSH_KEY", sshPrivateKeyPath)

	tasks := strings.Fields(vargs.Tasks)

	if len(tasks) == 0 {
		fmt.Println("Please provide Capistrano tasks to execute")
		os.Exit(1)
		return
	}

	bundle := command(workspace, "bundle", "install", "--path", "build/bundle")
	bundle.Run()

	bundler_args := append([]string{"exec", "cap"}, tasks...)
	capistrano := command(workspace, "bundle", bundler_args...)
	if err := capistrano.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func command(w drone.Workspace, cmd string, args ...string) *exec.Cmd {
	c := exec.Command(cmd, args...)
	c.Dir = w.Path
	c.Env = os.Environ()
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}
