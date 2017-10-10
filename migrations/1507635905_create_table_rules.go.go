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
			tx.CreateTable(&rule.Rule{})
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Write your migration rollback script here
			tx.DropTableIfExists(&rule.Rule{})
			return nil
		},
	}
)
