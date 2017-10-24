package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for 1 second to finish processing")
		fmt.Println("Closing db connection")
		db.Engine.Close()
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()

	return nil
}
