package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/manyminds/api2go/jsonapi"
	uuid "github.com/satori/go.uuid"
)

type Workflow struct {
	ID              []byte           `gorm:"type:binary(16);primary_key" json:"-"`
	UserID          uint             `gorm:"unsigned;unique_index:workflows_name_user_id;index;not null" json:"-"`
	Name            string           `gorm:"unique_index:workflows_name_user_id;not null" json:"name" json:"name"`
	IsShared        bool             `gorm:"not null;default:0" json:"is_shared"`
	IsActivated     bool             `gorm:"not null;default:0" json:"is_activated"`
	Actions         []Action         `gorm:"many2many:workflow_actions;" json:"-"`
	WorkflowObjects []WorkflowObject `json:"-"`
	Rules           []Rule           `json:"-"`
	CreatedAt       time.Time        `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time        `gorm:"default:current_timestamp on update current_timestamp" json:"updated_at"`
	DeletedAt       *time.Time       `json:"deleted_at,omitempty"`
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
	return nil
}

func (workflow *Workflow) GetReferencedIDs() []jsonapi.ReferenceID {
	refs := make(
		[]jsonapi.ReferenceID,
		GetRefsCount(
			workflow.Actions,
			workflow.WorkflowObjects,
			workflow.Rules,
		))
	idx := 0
	for _, d := range workflow.Actions {
		refs[idx] = jsonapi.ReferenceID{
			ID:   d.GetID(),
			Type: "actions",
			Name: "actions",
		}
		idx++
	}

	for _, d := range workflow.WorkflowObjects {
		refs[idx] = jsonapi.ReferenceID{
			ID:   d.GetID(),
			Type: "workflow-objects",
			Name: "workflow-objects",
		}
		idx++
	}

	for _, d := range workflow.Rules {
		refs[idx] = jsonapi.ReferenceID{
			ID:   d.GetID(),
			Type: "rules",
			Name: "rules",
		}
		idx++
	}

	return refs
}

func (workflow *Workflow) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	refs := make(
		[]jsonapi.MarshalIdentifier,
		GetRefsCount(workflow.Actions, workflow.WorkflowObjects, workflow.Rules),
	)
	idx := 0

	for _, d := range workflow.Actions {
		refs[idx] = d
		idx++
	}

	for _, d := range workflow.WorkflowObjects {
		refs[idx] = d
		idx++
	}

	for _, d := range workflow.Rules {
		refs[idx] = d
		idx++
	}

	return refs
}

func (w *Workflow) SetToOneReferenceID(name, ID string) error {
	return nil
}

func (w *Workflow) SetToManyReferenceIDs(name string, IDs []string) error {
	if len(IDs) <= 0 {
		return nil
	}

	if name == "rules" {
		w.Rules = make([]Rule, len(IDs))
		for i, ID := range IDs {
			w.Rules[i] = Rule{}
			w.Rules[i].SetID(ID)
			db.Engine.Find(&w.Rules[i])
		}
		return nil
	}
	return nil
}
