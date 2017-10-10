package rule

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Rule struct {
	ID         string    `sql:"type:varchar(36);primary_key"`
	WorkflowID string    `gorm:"varchar(36);index"`
	UserID     uint64    `gorm:"unsigned;unique_index:rules_name_user_id"`
	Name       string    `gorm:"not null;unique_index:rules_name_user_id"`
	CreatedAt  time.Time `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt  time.Time
}

func (r *Rule) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
