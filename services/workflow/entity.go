package workflow

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Workflow struct {
	ID          string    `sql:"type:varchar(36);primary_key"`
	UserID      uint64    `gorm:"unsigned;unique_index:workflows_name_user_id;index;not null"`
	Name        string    `gorm:"unique_index:workflows_name_user_id;not null"`
	IsShared    bool      `gorm:"not null;default:0"`
	IsActivated bool      `gorm:"not null;default:0"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt   *time.Time
}

func (w *Workflow) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
