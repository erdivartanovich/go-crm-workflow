package migrations

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/migrate"
	"github.com/kwri/go-workflow/services/entity"
)

var (
	create_table_tasks_1510558313 = migrate.Migration{
		ID: "1510558313",
		Migrate: func(tx *gorm.DB) error {
			err := tx.CreateTable(&entity.Task{}).Error
			if err != nil {
				return err
			}
			seedData := []*entity.Task{
				{
					UserID:            1,
					TaskType:          4,
					TaskAction:        "Four calls per year",
					DueDate:           time.Now(),
					FromInteraction:   "from interaction",
					Reason:            "Four calls per year",
					Description:       "Reach out and try to convert",
					IsCompleted:       1,
					IsAutomated:       0,
					CreatedBy:         1,
					UpdatedBy:         1,
					Status:            1,
					MinimumCompletion: 4,
				},
				{
					UserID:            1,
					TaskType:          4,
					TaskAction:        "Upcoming Birthday",
					DueDate:           time.Now(),
					FromInteraction:   "from interaction",
					Reason:            "Birthday",
					Description:       "Let them know you're thinking of them!",
					IsCompleted:       0,
					IsAutomated:       0,
					CreatedBy:         1,
					UpdatedBy:         1,
					Status:            1,
					MinimumCompletion: 9,
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
			err := tx.DropTableIfExists(&entity.Task{}).Error
			return err
		},
	}
)
