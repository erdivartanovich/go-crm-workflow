package rule

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Rule struct {
	ID         []byte    `gorm:"type:binary(16);primary_key"`
	WorkflowID []byte    `gorm:"type:binary(16);index"`
	UserID     uint64    `gorm:"unsigned;unique_index:rules_name_user_id"`
	Name       string    `gorm:"not null;unique_index:rules_name_user_id"`
	CreatedAt  time.Time `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt  time.Time
}

func (r *Rule) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4().MarshalBinary()
	scope.SetColumn("ID", uuid)
	return err
}

func (rule *Rule) GetKey() string {
	return string(rule.ID[:])
}
