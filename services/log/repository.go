package log

import (
	"errors"
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

func (repo *LogRepository) SetAdapter(adapter SearchAdapter) *LogRepository {
	repo.adapter = &adapter
	return repo
}

func (repo *LogRepository) Find() ([]*entity.WorkflowLog, error) {
	logs := &[]*entity.WorkflowLog{}
	err := repo.prepareDb().Find(logs).Error
	repo.ResetInstance()
	return *logs, err
}

func (repo *LogRepository) Where(workflowlog entity.WorkflowLog) *LogRepository {
	repo.where.Push(&workflowlog)
	return repo
}

func (repo *LogRepository) First() (*entity.WorkflowLog, error) {
	log := &entity.WorkflowLog{}
	err := repo.prepareDb().First(log).Error
	repo.ResetInstance()
	return log, err
}

func (repo *LogRepository) Update(workflowlog entity.WorkflowLog, payload entity.WorkflowLog) (*entity.WorkflowLog, error) {
	err := repo.prepareDb().Model(&workflowlog).Update(payload).Error
	repo.ResetInstance()
	return &workflowlog, err
}

func (repo *LogRepository) Replace(workflowlog entity.WorkflowLog, payload entity.WorkflowLog) (*entity.WorkflowLog, error) {
	wl := &workflowlog
	db := repo.prepareDb()
	db.First(wl)
	wl.ResourceName = payload.ResourceName
	err := db.Save(wl).Error
	repo.ResetInstance()
	return wl, err
}

func (repo *LogRepository) Insert(workflowlog entity.WorkflowLog) (*entity.WorkflowLog, error) {
	in := &workflowlog
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *LogRepository) Delete(workflowlog entity.WorkflowLog) (*entity.WorkflowLog, error) {
	in := &workflowlog
	if len(in.ID) == 0 {
		return nil, errors.New("You need to set ID of deleted workflow")
	}
	err := repo.prepareDb().Delete(&in).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Unscoped().Find(&in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *LogRepository) prepareDb() *gorm.DB {
	count := repo.where.Len()
	tx := repo.db
	for i := 0; i < count; i++ {
		tx = tx.Where(repo.where.Pop())
	}
	return tx
}

func (repo *LogRepository) ResetInstance() {
	repo.adapter = nil
}

func (repo *LogRepository) Count() (int, error) {
	count := 0
	err := repo.prepareDb().Model(&entity.WorkflowLog{}).Count(&count).Error
	return count, err
}

func NewLogRepository() *LogRepository {
	db := db.Engine
	return &LogRepository{
		db: db,
		where: &stack.Stack{},
	}
}