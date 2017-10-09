package migrator

import (
	"strconv"
	"strings"
	"time"

	m "github.com/kwri/go-workflow/migrations"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/modules/fs"
	"github.com/kwri/go-workflow/modules/logger"
	"github.com/kwri/go-workflow/modules/migrate"
)

func Migrate() error {
	engine := db.Engine
	defer engine.Close()

	if err := engine.DB().Ping(); err != nil {
		logger.Fatal(err)
	}

	m := migrate.New(engine, migrate.DefaultOptions, m.Migrations)
	return m.Migrate()
}

func RollBack() error {
	engine := db.Engine
	defer engine.Close()
	var err error
	if err = engine.DB().Ping(); err != nil {
		logger.Fatal(err)
	}

	m := migrate.New(engine, migrate.DefaultOptions, m.Migrations)
	return m.Rollback()
}

func Create(name string) error {
	id := generateMigrationID()
	filename := generateSqlFilename(name, id)
	createMigrationScript(filename, id)

	return nil
}

func generateMigrationID() string {
	t := time.Now().Unix()
	return strconv.FormatInt(t, 10)
}

func generateSqlFilename(name string, id string) string {
	return name + "_" + id
}

func createMigrationScript(name string, id string) {
	content := strings.Replace(template, "{{name}}", name, 1)
	content = strings.Replace(content, "{{id}}", id, 1)
	fs.FilePutContent("./migrations/"+name+".go", content)
}

var template = `package migrations

import (
    "github.com/jinzhu/gorm"
    "github.com/kwri/go-workflow/modules/migrate"

)

var (
	{{name}} = migrate.Migration{
        ID: "{{id}}",
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
