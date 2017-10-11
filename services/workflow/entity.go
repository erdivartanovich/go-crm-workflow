package workflow

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Workflow struct {
	ID          []byte     `gorm:"type:binary(16);primary_key" json:"id"`
	UserID      uint64     `gorm:"unsigned;unique_index:workflows_name_user_id;index;not null"  jsonapi:"attr,user_id"`
	Name        string     `gorm:"unique_index:workflows_name_user_id;not null"  jsonapi:"attr,name"`
	IsShared    bool       `gorm:"not null;default:0" jsonapi:"attr,is_shared"`
	IsActivated bool       `gorm:"not null;default:0" jsonapi:"attr,is_activated"`
	CreatedAt   time.Time  `gorm:"default:current_timestamp" jsonapi:"attr,created_at"`
	UpdatedAt   time.Time  `gorm:"default:current_timestamp on update current_timestamp" jsonapi:"attr,updated_at"`
	DeletedAt   *time.Time `jsonapi:"attr,deleted_at"`
}

func (w *Workflow) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4().MarshalBinary()
	scope.SetColumn("ID", uuid)
	return err
}

func (workflow *Workflow) GetKey() string {
	return string(workflow.ID[:])
}

func (workflow *Workflow) MarshalJSON() ([]byte, error) {
	id := &uuid.UUID{}
	copy(id[:], workflow.ID)
	type Alias Workflow
	return json.Marshal(&struct {
		ID string `json:"id"`
		*Alias
	}{
		ID:    id.String(),
		Alias: (*Alias)(workflow),
	})
}

func (workflow *Workflow) UnmarshalJSON() ([]byte, error) {

}
