package cmd

import (
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/modules/setting"
	"gopkg.in/urfave/cli.v2"
)

var (
	configFlag = &cli.StringFlag{
		Name:  "config, c",
		Value: "config.ini",
		Usage: "Configuration file path",
	}
)

func defaultAction(ctx *cli.Context) error {
	setting.ConfigFile = ctx.String("config")
	setting.Initialize()
	err := db.Initialize()
	if err != nil {
		return err
	}
	return nil
}
