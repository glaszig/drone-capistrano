package main

import (
	"fmt"
	"os"
	"log"

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
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{Tasks: c.String("tasks")}
	return plugin.Exec()
}
