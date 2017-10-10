package action

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Action struct {
	ID          string    `sql:"type:varchar(36);primary_key"`
	UserID      uint64    `gorm:"unsigned user_id;unique_index:actions_name_user_id;index"`
	TaskID      string    `sql:"type:varchar(36);index"`
	Name        string    `gorm:"not null;unique_index:actions_name_user_id"`
	ActionType  int8      `gorm:"not null;default:0"`
	TargetClass string    `gorm:"not null"`
	TargetField string    `gorm:"not null"`
	Value       string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt   *time.Time
}

func (r *Action) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

func (action *Action) GetValue() map[interface{}]interface{} {
	v := make(map[interface{}]interface{}, 0)
	data := []byte(action.Value)
	json.Unmarshal(data, &v)
	return v
}
