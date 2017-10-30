package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/log"
)

var (
	create_workflow_log_table_1509349752 = migrate.Migration{
		ID: "1509349752",
		Migrate: func(tx *gorm.DB) error {
			err := tx.CreateTable(&log.WorkflowLog{}).Error
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.DropTableIfExists(&log.WorkflowLog{}).Error
			return err
		},
	}
)
