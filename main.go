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

	fmt.Printf("Starting SSH agent\n")
	command(workspace, "eval `ssh-agent -s`", sshPrivateKeyPath)

	fmt.Printf("Adding deploy key to SSH agent\n")
	command(workspace, "ssh-add", sshPrivateKeyPath)

	tasks := strings.Fields(vargs.Tasks)

	if len(tasks) == 0 {
		fmt.Println("Please provide Capistrano tasks to execute")
		os.Exit(1)
		return
	}

	bundle := exec.Command("bundle", "install")
	bundle.Dir = workspace.Path
	bundle.Stderr = os.Stderr
	bundle.Stdout = os.Stdout
	bundle.Run()

	capistrano := exec.Command("bundle exec cap", tasks...)

	capistrano.Dir = workspace.Path
	capistrano.Stderr = os.Stderr
	capistrano.Stdout = os.Stdout

	if err := capistrano.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func command(w drone.Workspace, cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	c.Dir = w.Path
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}
