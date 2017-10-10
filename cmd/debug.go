package cmd

import (
	"github.com/kwri/go-workflow/modules/logger"
	"gopkg.in/urfave/cli.v2"
)

var (
	CmdDebug = cli.Command{
		Name:        "debug",
		Usage:       "Start workflow api web server",
		Description: "Run workflow rest api server",
		Action:      runDebug,
		Flags: []cli.Flag{
			configFlag,
		},
	}
)

func runDebug(ctx *cli.Context) error {
	logger.Initialize("debug")
	err := defaultAction(ctx)
	if err != nil {
		return err
	}

	debug()
	return nil

}

func debug() {

}
