package migrate

import (
	"bytes"

	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/logger"
)

var (
	DefaultOptions = Options{
		TableName: "migrations",
	}
)

type Migration struct {
	ID       string
	Migrate  MigrateFunc
	Rollback RollbackFunc
}

type Migratorable interface {
	Migrate() error
	Rollback() error
}

type Migrator struct {
	Engine     *gorm.DB
	Migrations []*Migration
	Options    Options
}

type Options struct {
	TableName string
}

type MigrateFunc func(tx *gorm.DB) error

type RollbackFunc func(tx *gorm.DB) error

//New create new migratorable instance
func New(tx *gorm.DB, options Options, migrations []*Migration) Migratorable {
	return Migrator{
		Engine:     tx,
		Migrations: migrations,
		Options:    options,
	}
}

func (migrator Migrator) Migrate() error {
	lastRunningMigration := migrator.getLastMigrationID()
	begin := false
	var err error
	var migrated []interface{}

	for _, m := range migrator.Migrations {

		if lastRunningMigration == "" {
			begin = true
		}
		if m.ID == lastRunningMigration {
			begin = true
			continue
		}
		if !begin {
			continue
		}
		logger.Infof("Migrating migration file with id: %s", m.ID)
		err = m.Migrate(migrator.Engine)
		if err != nil {
			break
		}
		migrated = append(migrated, []string{m.ID})
		logger.Infof("Success run migration file with id: %s", m.ID)
		lastRunningMigration = m.ID
	}
	migratedLen := len(migrated)
	targetLoopLen := migratedLen - 1
	if migratedLen > 0 {
		var insertIdsStmt bytes.Buffer
		insertIdsStmt.WriteString("insert into ")
		insertIdsStmt.WriteString(migrator.Options.TableName)
		insertIdsStmt.WriteString(" (id) VALUES")
		for idx := range migrated {
			insertIdsStmt.WriteString("(?)")
			if targetLoopLen != idx {
				insertIdsStmt.WriteString(",")
			}
		}

		migrator.Engine.Exec(insertIdsStmt.String(), migrated...)
	}

	return err
}

func (migrator Migrator) Rollback() error {
	lastMigrationID := migrator.getLastMigrationID()
	count := len(migrator.Migrations)
	begin := false
	var err error

	for i := count - 1; i >= 0; i-- {

		if lastMigrationID == "" {
			break
		}
		m := migrator.Migrations[i]
		if lastMigrationID == m.ID {
			begin = true
		}

		if begin {
			err = m.Rollback(migrator.Engine)

			if err != nil {

				break
			}
			migrator.Engine.Exec(
				"DELETE FROM "+migrator.Options.TableName+" where id = ?",
				m.ID,
			)
		}
	}
	if err == nil {
		logger.Info("All rollbacks run successfully.")
	}
	return err
}

func (migrator Migrator) getLastMigrationID() string {
	var result = &struct {
		ID string `xorm:"id"`
	}{}
	migrator.Engine.Raw(
		"select id from " + migrator.Options.TableName).
		Limit(1).
		Order("id desc").
		Scan(result)

	return result.ID
}
