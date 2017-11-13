package workflow

import (
	"errors"

	"github.com/golang-collections/collections/stack"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/services/entity"
)

type WorkflowRepostitory struct {
	db      *gorm.DB
	adapter *entity.SearchAdapter
	where   *stack.Stack
}

func (repo *WorkflowRepostitory) SetAdapter(adapter *entity.SearchAdapter) *WorkflowRepostitory {
	repo.adapter = adapter
	return repo
}

func (repo *WorkflowRepostitory) Find() ([]*entity.Workflow, error) {
	workflows := &[]*entity.Workflow{}
	err := repo.prepareDb().Find(workflows).Error
	repo.ResetInstance()
	return *workflows, err
}

func (repo *WorkflowRepostitory) Where(workflow entity.Workflow) *WorkflowRepostitory {
	repo.where.Push(&workflow)
	return repo
}

func (repo *WorkflowRepostitory) First() (*entity.Workflow, error) {

	workflow := &entity.Workflow{}
	err := repo.prepareDb().First(workflow).Error
	repo.ResetInstance()
	return workflow, err
}

func (repo *WorkflowRepostitory) Update(workflow entity.Workflow, payload entity.Workflow) (*entity.Workflow, error) {
	payload.ID = workflow.ID
	err := repo.prepareDb().Model(&workflow).Update(payload).Error
	repo.ResetInstance()
	return &workflow, err
}

func (repo *WorkflowRepostitory) Replace(workflow entity.Workflow, payload entity.Workflow) (*entity.Workflow, error) {
	wk := &workflow
	db := repo.prepareDb()
	db.First(wk)
	wk.IsActivated = payload.IsActivated
	wk.IsShared = payload.IsShared
	wk.Name = payload.Name
	err := db.Save(wk).Error
	repo.ResetInstance()
	return wk, err
}

func (repo *WorkflowRepostitory) Insert(workflow entity.Workflow) (*entity.Workflow, error) {

	in := &workflow
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *WorkflowRepostitory) Delete(workflow entity.Workflow) (*entity.Workflow, error) {

	in := &workflow
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

func (repo *WorkflowRepostitory) prepareDb() *gorm.DB {
	count := repo.where.Len()
	tx := repo.db
	for i := 0; i < count; i++ {
		tx = tx.Where(repo.where.Pop())
	}
	if repo.adapter != nil {
		tx = repo.adapter.ApplySearchAdapter(tx)
	}
	return tx
}

func (repo *WorkflowRepostitory) ResetInstance() {
	repo.adapter = nil
}

func (repo *WorkflowRepostitory) Count() (int, error) {
	count := 0
	tx := repo.prepareDb().Limit(-1).Offset(-1)
	err := tx.Model(&entity.Workflow{}).Count(&count).Error
	return count, err
}

func NewWorkflowRepository() *WorkflowRepostitory {
	db := db.Engine
	return &WorkflowRepostitory{
		db:    db,
		where: &stack.Stack{},
	}
}
