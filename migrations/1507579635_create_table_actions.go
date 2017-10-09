package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
)

var (
	create_table_actions_1507579635 = migrate.Migration{
		ID: "1507579635",
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
