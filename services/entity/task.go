package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Task struct {
	ID                  []byte     `gorm:"type:binary(16);primary_key" json:"-"`
	UserID              uint       `gorm:"unsigned;index" json:"-"`
	TaskType            int        `gorm:"type:smallint(6);" json:"-"`
	TaskAction          string     `json:"task_action"`
	DueDate             time.Time  `json:"due_date"`
	FromInteraction     string     `json:"from_interaction"`
	Reason              string     `json:"reason"`
	Description         string     `json:"description"`
	IsCompleted         int        `gorm:"type:tinyint(1);default:0" json:"is_completed"`
	IsAutomated         int        `gorm:"type:tinyint(1);default:0" json:"is_automated"`
	CreatedBy           uint       `gorm:"unsigned;index" json:"-"`
	UpdatedBy           uint       `gorm:"unsigned;index" json:"-"`
	CreatedAt           time.Time  `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt           time.Time  `gorm:"default:current_timestamp on update current_timestamp" json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at,omitempty"`
	Status              int        `gorm:"type:smallint(6)" json:"status"`
	PermanenttDeletedAt time.Time  `json:"permanent_deleted_at,omitempty"`
	MinimumCompletion   uint       `gorm:"unsigned;index" json:"-"`
}

func (r *Task) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4().MarshalBinary()
	scope.SetColumn("ID", uuid)
	return err
}

func (rule *Task) GetID() string {
	id := &uuid.UUID{}
	copy(id[:], rule.ID)
	return id.String()
}

func (rule *Task) SetID(id string) error {
	rule.UnmarshalUUIDString(id)
	return nil
}

func (rule *Task) UnmarshalUUIDString(id string) {
	uuid := &uuid.UUID{}
	uuid.UnmarshalText([]byte(id))
	binid, _ := uuid.MarshalBinary()
	rule.ID = binid
}
