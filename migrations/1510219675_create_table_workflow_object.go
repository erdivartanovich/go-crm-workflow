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
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.DropTableIfExists(&entity.WorkflowObject{}).Error
			return err
		},
	}
)
