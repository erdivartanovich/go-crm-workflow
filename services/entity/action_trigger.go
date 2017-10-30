package entity

import (
	"time"
	"github.com/google/jsonapi"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type ActionTrigger struct {
	ID          []byte    `gorm:"type:binary(16);primary_key" json:"-"`
	ActionID   	uint64    `gorm:"type:binary(16);index" json:"-"`
	TargetField string    `gorm:"not null" json:"target_field"`
	Min			string	  `gorm:"not null;default:'*'" json:"min"`
	Hour		string	  `gorm:"not null;default:'*'" json:"hour"`
	DayPerMonth string	  `gorm:"not null;default:'*'" json:"day_per_month"`
	Month		string	  `gorm:"not null;default:'*'" json:"month"`
	DayPerWeek 	string	  `gorm:"not null;default:'*'" json:"day_per_week"`
	RunnableOnce string	  `gorm:"not null;default:'*'" json:"runnable_once"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp on update current_timestamp" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
} 

func (a *ActionTrigger) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4().MarshalBinary()
	scope.SetColumn("ID", uuid)
	return err
}

func (actiontrigger *ActionTrigger) GetID() string {
	id := &uuid.UUID{}
	copy(id[:], actiontrigger.ID)
	return id.String()
}

func (actiontrigger *ActionTrigger) UnmarshalUUIDString(id string) {
	uuid := &uuid.UUID{}
	uuid.UnmarshalText([]byte(id))
	binid, _ := uuid.MarshalBinary()
	actiontrigger.ID = binid
}

func (actiontrigger *ActionTrigger) GetCustomLinks(link string) jsonapi.Links {
	links := make(jsonapi.Links)
	links["current"] = link
	return links
}