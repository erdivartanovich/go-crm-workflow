package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type WorkflowLog struct {
	ID           []byte    `gorm:"type:binary(16);primary_key" json:"-"`
	UserID       uint      `gorm:"unsigned"`
	ActionID     []byte    `gorm:"type:binary(16)" json:"-"`
	WorkflowID   []byte    `gorm:"type:binary(16)" json:"-"`
	ResourceID   uint      `gorm:"unisgned nullable" json:"-"`
	ResourceName string    `gorm:"nullable"`
	Status       int       `gorm:"type:tinyint(1)"`
	Info         string    `gorm:"type:text"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt    *time.Time
}

func (log *WorkflowLog) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4().MarshalBinary()
	scope.SetColumn("ID", uuid)
	return err
}

func (log *WorkflowLog) GetKey() string {
	return string(log.ID[:])
}
