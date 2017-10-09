package cmd

import (
	"github.com/kwri/go-workflow/modules/logger"
	"gopkg.in/urfave/cli.v2"
)

var (
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "port, p",
			Value: "3000",
			Usage: "Temporary port number to prevent conflict",
		},
		configFlag,
		&cli.StringFlag{
			Name:  "pid, P",
			Value: "/var/run/workflow.pid",
			Usage: "Custom pid file path",
		},
	}

	CmdApi = cli.Command{
		Name:        "api",
		Usage:       "Start workflow api web server",
		Description: "Run workflow rest api server",
		Action:      runApiService,
		Flags:       flags,
	}
)

func runApiService(ctx *cli.Context) error {
	logger.Initialize("api")
	err := defaultAction(ctx)
	return err
}
