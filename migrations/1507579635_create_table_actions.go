package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/entity"
	"github.com/kwri/go-workflow/services/entity/constant"
)

var (
	create_table_actions_1507579635 = migrate.Migration{
		ID: "1507579635",
		Migrate: func(tx *gorm.DB) error {
			err := tx.CreateTable(&entity.Action{}).Error
			if err != nil {
				return err
			}
			workflow := entity.Workflow{}
			tx.First(&workflow)
			workflows := []entity.Workflow{workflow}
			seedData := []*entity.Action{
				{
					Name:        "Mailchimp Discount Campaign",
					TaskID:      "",
					UserID:      1,
					ActionType:  constant.ACTION_TYPE_MAILCHIMP,
					TargetClass: "",
					TargetField: "",
					Value:       "",
					Workflows:   workflows,
				},
				{
					Name:        "Mailchimp Birthday Bonus Campaign",
					TaskID:      "",
					UserID:      1,
					ActionType:  constant.ACTION_TYPE_MAILCHIMP,
					TargetClass: "",
					TargetField: "",
					Value:       "",
					Workflows:   workflows,
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
			err := tx.DropTableIfExists(&entity.Action{}).Error
			return err
		},
	}
)
