package cmd

import (
	"errors"

	"github.com/kwri/go-workflow/modules/generator"
	"github.com/kwri/go-workflow/modules/logger"
	"gopkg.in/urfave/cli.v2"
)

var (
	CmdMakeService = cli.Command{
		Name:   "make:service",
		Usage:  "Create service boilerplate",
		Action: runMakeServiceService,
		Flags: []cli.Flag{
			configFlag,
		},
	}
)

func runMakeServiceService(ctx *cli.Context) error {
	logger.Initialize("generator")
	err := defaultAction(ctx)
	if err != nil {
		return err
	}
	name := ctx.Args().First()
	if name == "" {
		return errors.New("No name suplied.")
	}
	generator.NewService(name).Generate()
	return nil
}
