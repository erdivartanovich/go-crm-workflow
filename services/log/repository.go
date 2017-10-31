package log

import (
	"github.com/golang-collections/collections/stack"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/services/entity"
)

type LogRepository struct {
	db      *gorm.DB
	adapter *SearchAdapter
	where   *stack.Stack
}

func NewLogRepository() *LogRepository {
	db := db.Engine
	return &LogRepository{
		db: db,
		where: &stack.Stack{},
	}
}