package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var build string = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "Capistrano plugin"
	app.Usage = "Capistrano plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "tasks",
			Usage:  "capistrano tasks to run",
			EnvVar: "PLUGIN_TASKS",
		},
		cli.StringFlag{
			Name:   "private_key",
			Usage:  "SSH private key",
			EnvVar: "CAPISTRANO_PRIVATE_KEY",
		},
		cli.StringFlag{
			Name:   "public_key",
			Usage:  "SSH public key",
			EnvVar: "CAPISTRANO_PUBLIC_KEY",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Config: Config{
			Tasks:      c.String("tasks"),
			PrivateKey: c.String("private_key"),
			PublicKey:  c.String("public_key"),
		},
	}
	return plugin.Exec()
}
