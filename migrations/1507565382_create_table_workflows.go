package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/workflow"
)

var (
	create_table_workflows_1507565382 = migrate.Migration{
		ID: "1507565382",
		Migrate: func(tx *gorm.DB) error {
			tx.CreateTable(&workflow.Workflow{})
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			tx.DropTableIfExists(&workflow.Workflow{})
			return nil
		},
	}
)
