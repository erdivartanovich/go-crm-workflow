package migrations

import (
	"github.com/kwri/go-workflow/services/entity"
    "github.com/jinzhu/gorm"
    "github.com/kwri/go-workflow/modules/migrate"
)

var (
	create_table_action_trigger_1509351627 = migrate.Migration{
        ID: "1509351627",
        Migrate: func(tx *gorm.DB) error {
            err := tx.CreateTable(&entity.ActionTrigger{}).Error
            return err
        },
        Rollback: func(tx *gorm.DB) error {
            err := tx.DropTableIfExists(&entity.ActionTrigger{}).Error
            return err
        },
    }
)

