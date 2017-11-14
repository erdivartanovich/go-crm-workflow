package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/entity"
)

var (
	add_rule_seeder_1510572541 = migrate.Migration{
		ID: "1510572541",
		Migrate: func(tx *gorm.DB) error {
			workflow := entity.Workflow{}
			tx.First(&workflow)
			action := entity.Action{}
			tx.First(&action)
			actions := []entity.Action{action}
			seedData := []*entity.Rule{
				{
					WorkflowID: workflow.ID,
					Name:       "born_in_december",
					RuleType:   6,
					FieldName:  "persons.date_of_birth",
					Operator:   1,
					Value:      "12",
					Priority:   2,
					UserID:     1,
					Actions:    actions,
				},
				{
					WorkflowID: workflow.ID,
					Name:       "born_in_january",
					RuleType:   6,
					FieldName:  "persons.date_of_birth",
					Operator:   1,
					Value:      "1",
					Priority:   1,
					UserID:     1,
					Actions:    actions,
				},
			}
			db := tx.Begin()
			for _, data := range seedData {
				err := db.Create(data).Error
				if err != nil {
					db.Rollback()
					return err
				}
			}

			db.Commit()
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Write your migration rollback script here
			return nil
		},
	}
)
