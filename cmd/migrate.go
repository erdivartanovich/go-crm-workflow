package cmd

import (
	"errors"
	"sync"

	"github.com/kwri/go-workflow/modules/logger"
	"github.com/kwri/go-workflow/modules/migrator"
	"gopkg.in/urfave/cli.v2"
)

var (
	migrateFlags = []cli.Flag{
		configFlag,
	}
	CmdMigrate = cli.Command{
		Name:  "migrate",
		Usage: "Run db migration",
		Subcommands: []*cli.Command{
			&cli.Command{
				Name:        "migrate",
				Usage:       "Run db migrations",
				Description: "Run db migrations",
				Action:      actionMigrate,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "refresh",
					},
					configFlag,
				},
			},
			&cli.Command{
				Name:        "rollback",
				Usage:       "Rollback db migrations",
				Description: "Rollback db migrations",
				Action:      actionRollbackMigration,
				Flags:       migrateFlags,
			},
			&cli.Command{
				Name:        "create",
				Usage:       "create [script name]",
				Description: "create db migrations",
				Action:      actionCreateMigration,
				Flags: []cli.Flag{
					configFlag,
				},
			},
		},
	}
)

func actionMigrate(ctx *cli.Context) error {
	logger.Initialize("migrate")
	err := defaultAction(ctx)

	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	ch := make(chan error)
	if ctx.Bool("refresh") {
		wg.Add(1)
		go func() {
			wg.Done()
			ch <- migrator.RollBack()
		}()
		wg.Wait()
		err = <-ch
		if err != nil {
			return err
		}
	}

	return migrator.Migrate()

}

func actionRollbackMigration(ctx *cli.Context) error {
	logger.Initialize("migrate")
	err := defaultAction(ctx)
	if err != nil {
		return err
	}
	return migrator.RollBack()
}

func actionCreateMigration(ctx *cli.Context) error {
	logger.Initialize("migrate")
	err := defaultAction(ctx)
	if err != nil {
		return err
	}
	name := ctx.Args().First()
	if name == "" {
		return errors.New("No name suplied.")
	}
	return migrator.Create(name)
}
