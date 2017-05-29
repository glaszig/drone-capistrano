package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type (
	Config struct {
		Tasks      string
		PrivateKey string
		PublicKey  string
	}

	Plugin struct {
		Config Config
	}

	DeployWorkspace struct{}
)

// Exec executes the plugin step
func (p Plugin) Exec() error {
	writeSshKey(p.Config)

	dw := DeployWorkspace{}
	tasks := strings.Fields(p.Config.Tasks)

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
	return w.command("/bundle.sh", args...)
}

func (w *DeployWorkspace) command(cmd string, args ...string) *exec.Cmd {
	c := exec.Command(cmd, args...)
	// c.Dir = w.Workspace.Path
	c.Env = os.Environ()
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}

func writeSshKey(c Config) error {
	var err error = nil
	var private_key_bytes []byte = nil
	var public_key_bytes []byte = nil

	private_key_bytes, err = base64.StdEncoding.DecodeString(c.PrivateKey)
	if err != nil {
		return fmt.Errorf("Failed decoding private key: %s", err)
	}

	err = ioutil.WriteFile("/root/.ssh/capistrano", private_key_bytes, 0600)
	if err != nil {
		return fmt.Errorf("Failed writing private key: %s", err)
	}

	public_key_bytes, err = base64.StdEncoding.DecodeString(c.PublicKey)
	if err != nil {
		return fmt.Errorf("Failed decoding public key: %s", err)
	}

	err = ioutil.WriteFile("/root/.ssh/capistrano.pub", public_key_bytes, 0644)
	if err != nil {
		return fmt.Errorf("Failed writing public key: %s", err)
	}

	return nil
}

func printLog(message string, a ...interface{}) {
	fmt.Printf("=> %s\n", fmt.Sprintf(message, a...))
}
