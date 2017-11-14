package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/entity"
)

var (
	create_table_action_trigger_1509351627 = migrate.Migration{
		ID: "1509351627",
		Migrate: func(tx *gorm.DB) error {
			err := tx.CreateTable(&entity.ActionTrigger{}).Error
			if err != nil {
				return err
			}
			action := entity.Action{}
			tx.First(&action)
			seedData := []*entity.ActionTrigger{
				{
					ActionID:    action.ID,
					TargetField: "sendPrimary",
				},
				{
					ActionID:    action.ID,
					TargetField: "send",
				},
			}
			db := tx.Begin()
			for _, data := range seedData {
				err = db.Create(data).Error
				if err != nil {
					db.Rollback()
					return err
				}
			}

			db.Commit()
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.DropTableIfExists(&entity.ActionTrigger{}).Error
			return err
		},
	}
)
