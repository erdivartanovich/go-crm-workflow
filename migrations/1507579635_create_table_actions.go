package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/entity"
)

var (
	create_table_actions_1507579635 = migrate.Migration{
		ID: "1507579635",
		Migrate: func(tx *gorm.DB) error {
			err := tx.CreateTable(&entity.Action{}).Error
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.DropTableIfExists(&entity.Action{}).Error
			return err
		},
	}
)
