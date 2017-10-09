package cmd

import (
	"github.com/kwri/go-workflow/modules/logger"
	"gopkg.in/urfave/cli.v2"
)

var (
	CmdJob = cli.Command{
		Name:   "job",
		Usage:  "Run workflow job service daemon",
		Action: runJobService,
		Flags: []cli.Flag{
			configFlag,
		},
	}
)

func runJobService(ctx *cli.Context) error {
	logger.Initialize("job")
	err := defaultAction(ctx)
	return err
}
