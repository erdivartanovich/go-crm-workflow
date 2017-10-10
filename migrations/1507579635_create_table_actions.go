package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/action"
)

var (
	create_table_actions_1507579635 = migrate.Migration{
		ID: "1507579635",
		Migrate: func(tx *gorm.DB) error {
			tx.CreateTable(&action.Action{})
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			tx.DropTableIfExists(&action.Action{})
			return nil
		},
	}
)
