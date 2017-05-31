package main

import (
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
)

// Exec executes the plugin step
func (p Plugin) Exec() error {
	var err error = nil

	err = writeSshKey(p.Config)
	if err != nil {
		return err
	}

	tasks := strings.Fields(p.Config.Tasks)

	if len(tasks) == 0 {
		return fmt.Errorf("Please provide Capistrano tasks to execute")
	}

	printLog("Running Bundler")
	bundle := bundle("install")
	if err := bundle.Run(); err != nil {
		return fmt.Errorf("Bundler failed. %s", err)
	}

	printLog("Running Capistrano")
	capistrano := capistrano(tasks...)
	if err := capistrano.Run(); err != nil {
		return fmt.Errorf("Capistrano failed. %s", err)
	}

	return nil
}

func capistrano(tasks ...string) *exec.Cmd {
	args := append([]string{"exec", "cap"}, tasks...)
	return bundle(args...)
}

func bundle(args ...string) *exec.Cmd {
	return shellCommand("/bundle.sh", args...)
}

func shellCommand(cmd string, args ...string) *exec.Cmd {
	c := exec.Command(cmd, args...)
	// c.Dir = w.Workspace.Path
	c.Env = os.Environ()
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}

func writeSshKey(c Config) error {
	var err error = nil
	var private_key_bytes = []byte(c.PrivateKey)
	var public_key_bytes = []byte(c.PublicKey)

	_ = os.MkdirAll("/root/.ssh", 0700)

	err = ioutil.WriteFile("/root/.ssh/capistrano", private_key_bytes, 0600)
	if err != nil {
		return fmt.Errorf("Failed writing private key: %s", err)
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
