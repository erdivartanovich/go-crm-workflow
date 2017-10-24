package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/rule"
)

var (
	create_table_rules_1507635905 = migrate.Migration{
		ID: "1507635905",
		Migrate: func(tx *gorm.DB) error {
			// Write your migration script here
			err := tx.CreateTable(&rule.Rule{}).Error
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			// Write your migration rollback script here
			err := tx.DropTableIfExists(&rule.Rule{}).Error
			return err
		},
	}
)
