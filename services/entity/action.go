package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Action struct {
	ID          []byte     `gorm:"type:binary(16);primary_key" json:"-"`
	UserID      uint       `gorm:"unsigned user_id;unique_index:actions_name_user_id;index;not null" json:"-"`
	TaskID      string     `sql:"type:varchar(36);index" json:"task_id"`
	Name        string     `gorm:"not null;unique_index:actions_name_user_id" json:"name"`
	ActionType  int8       `gorm:"not null;default:0" json:"action_type"`
	TargetClass string     `gorm:"not null" json:"target_class"`
	TargetField string     `gorm:"not null" json:"target_field"`
	Value       string     `gorm:"type:text" json:"value"`
	Workflows   []Workflow `gorm:"many2many:workflow_actions;" json:"-"`
	Rules       []Rule     `gorm:"many2many:rule_actions;" json:"-"`
	CreatedAt   time.Time  `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:current_timestamp on update current_timestamp" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func (r *Action) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4().MarshalBinary()
	scope.SetColumn("ID", uuid)
	return err
}

func (action Action) GetID() string {
	id := &uuid.UUID{}
	copy(id[:], action.ID)
	return id.String()
}

func (action *Action) SetID(id string) error {
	action.UnmarshalUUIDString(id)
	return nil
}

func (action *Action) UnmarshalUUIDString(id string) {
	uuid := &uuid.UUID{}
	uuid.UnmarshalText([]byte(id))
	binid, _ := uuid.MarshalBinary()
	action.ID = binid
}
