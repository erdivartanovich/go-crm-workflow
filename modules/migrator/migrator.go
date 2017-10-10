package migrator

import (
	"fmt"
	"strconv"
	"time"

	m "github.com/kwri/go-workflow/migrations"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/modules/fs"
	"github.com/kwri/go-workflow/modules/migrate"
)

func Migrate() error {
	engine, err := db.NewEngine()
	if err != nil {
		return err
	}
	defer engine.Close()

	if err = engine.DB().Ping(); err != nil {
		return err
	}

	m := migrate.New(engine, migrate.DefaultOptions, m.Migrations)
	return m.Migrate()
}

func RollBack() error {
	engine, err := db.NewEngine()

	if err != nil {
		return err
	}
	defer engine.Close()

	if err = engine.DB().Ping(); err != nil {
		return err
	}

	m := migrate.New(engine, migrate.DefaultOptions, m.Migrations)
	return m.Rollback()
}

func Create(name string) error {
	id := generateMigrationID()
	createMigrationScript(name, id)

	return nil
}

func generateMigrationID() string {
	t := time.Now().Unix()
	return strconv.FormatInt(t, 10)
}

func generateSqlFilename(name string, id string) string {
	return "./migrations/" + id + "_" + name + ".go"
}

func createMigrationScript(name string, id string) {
	content := fmt.Sprintf(template, name+"_"+id, id)
	filename := generateSqlFilename(name, id)
	fs.FilePutContent(filename, content)
}

var template = `package migrations

import (
    "github.com/jinzhu/gorm"
    "github.com/kwri/go-workflow/modules/migrate"
)

var (
	%s = migrate.Migration{
        ID: "%s",
        Migrate: func(tx *gorm.DB) error {
            // Write your migration script here
            return nil
        },
        Rollback: func(tx *gorm.DB) error {
            // Write your migration rollback script here
            return nil
        },
    }
)

`
