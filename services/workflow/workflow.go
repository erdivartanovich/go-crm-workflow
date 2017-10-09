package workflow

import "time"

type Workflows struct {
	ID        uint64 `gorm:"primary_key"`
	UserID    uint64 `gorm:"unsigned user_id"`
	Name      string `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
