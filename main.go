package main

import (
	"os"

	log "github.com/kwri/go-workflow/modules/logger"

	"github.com/kwri/go-workflow/cmd"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	command := []*cli.Command{
		&cmd.CmdApi,
		&cmd.CmdJob,
		&cmd.CmdMigrate,
	}
	app := cli.App{
		Name:     "KW Workflow",
		Usage:    "Keller William's workflow api service & task runner",
		Commands: command,
		Version:  "1.0.0",
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal("Failed to run app with error", os.Args, err)
	}
}
