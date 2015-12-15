package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build     string
	buildDate string
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

	capistrano := exec.Command("cap", tasks...)

	capistrano.Dir = workspace.Path
	capistrano.Stderr = os.Stderr
	capistrano.Stdout = os.Stdout

	if err := capistrano.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}
