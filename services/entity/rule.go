package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Rule struct {
	ID         []byte     `gorm:"type:binary(16);primary_key" json:"-"`
	ParentID   []byte     `gorm:"type:binary(16);index" json:"-"`
	WorkflowID []byte     `gorm:"type:binary(16);index" json:"-"`
	UserID     uint       `gorm:"unsigned;unique_index:rules_name_user_id" json:"-"`
	Name       string     `gorm:"not null;unique_index:rules_name_user_id" json:"name"`
	RuleType   int        `gorm:"type:tinyint(1);not null" json:"rule_type"`
	FieldName  string     `gorm:"not null" json:"field_name"`
	Operator   int        `gorm:"type:tinyint(4);not null" json:"operator"`
	Value      string     `gorm:"not null" json:"value"`
	Priority   int        `gorm:"not null;type:tinyint(4)" json:"priority"`
	Actions    []Action   `gorm:"many2many:rule_action;" json:"-"`
	CreatedAt  time.Time  `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"default:current_timestamp on update current_timestamp" json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

func (r *Rule) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4().MarshalBinary()
	scope.SetColumn("ID", uuid)
	return err
}

func (rule *Rule) GetID() string {
	id := &uuid.UUID{}
	copy(id[:], rule.ID)
	return id.String()
}

func (rule *Rule) SetID(id string) error {
	rule.UnmarshalUUIDString(id)
	return nil
}

func (rule *Rule) UnmarshalUUIDString(id string) {
	uuid := &uuid.UUID{}
	uuid.UnmarshalText([]byte(id))
	binid, _ := uuid.MarshalBinary()
	rule.ID = binid
}
