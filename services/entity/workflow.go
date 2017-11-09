package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go/jsonapi"
	uuid "github.com/satori/go.uuid"
)

type Workflow struct {
	ID          []byte     `gorm:"type:binary(16);primary_key" json:"-"`
	UserID      uint       `gorm:"unsigned;unique_index:workflows_name_user_id;index;not null" json:"-"`
	Name        string     `gorm:"unique_index:workflows_name_user_id;not null" json:"name" json:"name"`
	IsShared    bool       `gorm:"not null;default:0" json:"is_shared"`
	IsActivated bool       `gorm:"not null;default:0" json:"is_activated"`
	Actions     []Action   `gorm:"many2many:workflow_actions;" json:"-"`
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

func (workflow *Workflow) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "action",
			Name: "actions",
		},
	}
}

func (workflow *Workflow) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, id := range workflow.Actions {
		result = append(result, jsonapi.ReferenceID{
			ID:   id.GetID(),
			Type: "action",
			Name: "actions",
		})
	}

	return result
}
