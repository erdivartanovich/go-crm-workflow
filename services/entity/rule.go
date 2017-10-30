package entity

import (
	"time"

	"github.com/google/jsonapi"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Rule struct {
	ID         []byte     `gorm:"type:binary(16);primary_key" json:"-"`
	WorkflowID []byte     `gorm:"type:binary(16);index" json:"-"`
	UserID     uint       `gorm:"unsigned;unique_index:rules_name_user_id" json:"-"`
	Name       string     `gorm:"not null;unique_index:rules_name_user_id" json:"name"`
	Actions    []Action   `gorm:"many2many:rule_action;" json:"actions"`
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

func (rule *Rule) UnmarshalUUIDString(id string) {
	uuid := &uuid.UUID{}
	uuid.UnmarshalText([]byte(id))
	binid, _ := uuid.MarshalBinary()
	rule.ID = binid
}

func (rule *Rule) GetCustomLinks(link string) jsonapi.Links {
	links := make(jsonapi.Links)
	links["current"] = link
	return links
}
