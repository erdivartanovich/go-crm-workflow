package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/entity"
)

var (
	create_table_rules_1507635905 = migrate.Migration{
		ID: "1507635905",
		Migrate: func(tx *gorm.DB) error {
			// Write your migration script here
			err := tx.CreateTable(&entity.Rule{}).Error
			if err != nil {
				return err
			}
			seedData := []*entity.Rule{
				{
					Name:      "Greater",
					RuleType:  1,
					FieldName: "Greater Than",
					Operator:  1,
					Value:     "1",
					Priority:  1,
					UserID:    1,
				},
				{
					Name:      "Lower",
					RuleType:  2,
					FieldName: "Lower Than",
					Operator:  2,
					Value:     "1",
					Priority:  1,
					UserID:    1,
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
			// Write your migration rollback script here
			err := tx.DropTableIfExists(&entity.Rule{}).Error
			return err
		},
	}
)
