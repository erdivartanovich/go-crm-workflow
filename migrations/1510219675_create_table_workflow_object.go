package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/entity"
)

var (
	create_table_workflow_object_1510219675 = migrate.Migration{
		ID: "1510219675",
		Migrate: func(tx *gorm.DB) error {
			err := tx.CreateTable(&entity.WorkflowObject{}).Error
			if err != nil {
				return err
			}
			workflow := entity.Workflow{}
			tx.First(&workflow)
			seedData := []*entity.WorkflowObject{
				{
					UserID:      1,
					WorkflowID:  workflow.ID,
					ObjectClass: "tags.id",
					ObjectType:  "1",
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
			err := tx.DropTableIfExists(&entity.WorkflowObject{}).Error
			return err
		},
	}
)
