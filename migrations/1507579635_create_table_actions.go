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
			if err != nil {
				return err
			}
			workflowSeeds1 := []entity.Workflow{
				{
					Name:   "Mailchimp New Year Discount Campaign",
					UserID: 1,
				},
				{
					Name:   "Mailchimp Holiday Bonus Campaign",
					UserID: 1,
				},
			}
			workflowSeeds2 := []entity.Workflow{
				{
					Name:   "Mailchimp Promotion Campaign",
					UserID: 1,
				},
				{
					Name:   "Mailchimp New Discount Campaign",
					UserID: 1,
				},
			}
			actionSeeds := []*entity.Action{
				{
					Name: "Action 1",
					UserID: 1,
					ActionType: 1,
					TaskID: "Task1",
					Workflows: workflowSeeds1,
				},
				{
					Name: "Action 2",
					UserID: 1,
					ActionType: 2,
					TaskID: "Task1",
					Workflows: workflowSeeds2,
				},
			}
			db := tx.Begin()
			for _, data := range actionSeeds {
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
