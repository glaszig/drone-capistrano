package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
)

type (
  // Plugin defines the Docker plugin parameters.
  Plugin struct {
    Tasks string
  }
  DeployWorkspace struct{}
)

// Exec executes the plugin step
func (p Plugin) Exec() error {
  dw := DeployWorkspace{}
  tasks := strings.Fields(p.Tasks)

  if len(tasks) == 0 {
    return fmt.Errorf("Please provide Capistrano tasks to execute")
  }

  printLog("Running Bundler")
  bundle := dw.bundle(bundlerArgs()...)
  if err := bundle.Run(); err != nil {
    return fmt.Errorf("Bundler failed. %s", err)
  }

  printLog("Running Capistrano")
  capistrano := dw.cap(tasks...)
  if err := capistrano.Run(); err != nil {
    return fmt.Errorf("Capistrano failed. %s", err)
  }

  return nil
}

func bundlerArgs() []string {
  args := []string{"install"}
  // if ! vargs.Debug {
  //   args = append(args, "--quiet")
  // }
  // if len(vargs.BundlePath) > 0 {
  //   args = append(args, "--path", vargs.BundlePath)
  // }
  return args
}

func (w *DeployWorkspace) cap(tasks ...string) *exec.Cmd {
  args := append([]string{"exec", "cap"}, tasks...)
  return w.bundle(args...)
}

func (w *DeployWorkspace) bundle(args ...string) *exec.Cmd {
  // return w.command("/bundle.sh", args...)
  return w.command("bundle", args...)
}

func (w *DeployWorkspace) command(cmd string, args ...string) *exec.Cmd {
  c := exec.Command(cmd, args...)
  // c.Dir = w.Workspace.Path
  c.Env = os.Environ()
  c.Stdout = os.Stdout
  c.Stderr = os.Stderr
  return c
}

func printLog(message string, a ...interface{}) {
  fmt.Printf("=> %s\n", fmt.Sprintf(message, a...))
}