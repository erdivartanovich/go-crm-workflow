package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/entity"
)

var (
	create_table_workflows_1507565382 = migrate.Migration{
		ID: "1507565382",
		Migrate: func(tx *gorm.DB) error {

			err := tx.CreateTable(&entity.Workflow{}).Error
			if err != nil {
				return err
			}
			seedData := []*entity.Workflow{
				{
					Name:   "Mailchimp Discount Campaign",
					UserID: 1,
				},
				{
					Name:   "Mailchimp Birthday Bonus Campaign",
					UserID: 1,
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
			err := tx.DropTableIfExists(&entity.Workflow{}).Error
			return err
		},
	}
)
