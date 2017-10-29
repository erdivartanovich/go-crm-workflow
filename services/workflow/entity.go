package workflow

import (
	"time"

	"github.com/google/jsonapi"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Workflow struct {
	ID          []byte     `gorm:"type:binary(16);primary_key" json:"-"`
	UserID      uint64     `gorm:"unsigned;unique_index:workflows_name_user_id;index;not null" json:"-"`
	Name        string     `gorm:"unique_index:workflows_name_user_id;not null" json:"name" json:"name"`
	IsShared    bool       `gorm:"not null;default:0" json:"is_shared"`
	IsActivated bool       `gorm:"not null;default:0" json:"is_activated"`
	CreatedAt   time.Time  `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:current_timestamp on update current_timestamp" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func (w *Workflow) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4().MarshalBinary()
	scope.SetColumn("ID", uuid)
	return err
}

func (workflow Workflow) GetID() string {
	id := &uuid.UUID{}
	copy(id[:], workflow.ID)
	return id.String()
}

func (workflow *Workflow) SetID(id string) error {
	workflow.UnmarshalUUIDString(id)
	return nil
}

func (workflow *Workflow) UnmarshalUUIDString(id string) error {
	uuid := &uuid.UUID{}
	uuid.UnmarshalText([]byte(id))
	binid, err := uuid.MarshalBinary()
	workflow.ID = binid
	return err
}

func (workflow *Workflow) GetCustomLinks(link string) jsonapi.Links {
	links := make(jsonapi.Links)
	links["current"] = link
	return links
}
