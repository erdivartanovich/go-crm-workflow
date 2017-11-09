package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go/jsonapi"
	uuid "github.com/satori/go.uuid"
)

type WorkflowObject struct {
	ID          []byte     `gorm:"type:binary(16);primary_key" json:"-"`
	UserID      uint       `gorm:"unsigned;index;not null" json:"-"`
	WorkflowID  []byte     `gorm:"type:binary(16);index;not null" json:"-"`
	ObjectClass string     `gorm:"not null" json:"object_class"`
	ObjectType  string     `gorm:"not null" json:"object_type"`
	Workflows   []Workflow `json:"-"`
	CreatedAt   time.Time  `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:current_timestamp on update current_timestamp" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func (o *WorkflowObject) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4().MarshalBinary()
	scope.SetColumn("ID", uuid)
	return err
}

func (o WorkflowObject) GetID() string {
	id := &uuid.UUID{}
	copy(id[:], o.ID)
	return id.String()
}

func (o *WorkflowObject) SetID(id string) error {
	o.UnmarshalUUIDString(id)
	return nil
}

func (o *WorkflowObject) UnmarshalUUIDString(id string) error {
	uuid := &uuid.UUID{}
	uuid.UnmarshalText([]byte(id))
	binid, err := uuid.MarshalBinary()
	o.ID = binid
	return err
}

func (o *WorkflowObject) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "workflow",
			Name: "workflows",
		},
	}
}

func (o *WorkflowObject) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, id := range o.Workflows {
		result = append(result, jsonapi.ReferenceID{
			ID:   id.GetID(),
			Type: "workflow",
			Name: "workflows",
		})
	}

	return result
}
