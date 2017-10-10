package workflow

import (
	"errors"

	"github.com/golang-collections/collections/stack"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
)

type WorkflowRepostitory struct {
	db      *gorm.DB
	adapter *SearchAdapter
	where   *stack.Stack
}

func (repo *WorkflowRepostitory) SetAdapter(adapter SearchAdapter) *WorkflowRepostitory {
	repo.adapter = &adapter
	return repo
}

func (repo *WorkflowRepostitory) Find() (*[]Workflow, error) {
	defer func() {
		repo.db.Close()
	}()
	workflows := &[]Workflow{}
	err := repo.prepareDb().Find(workflows).Error
	repo.ResetInstance()
	return workflows, err
}

func (repo *WorkflowRepostitory) Where(workflow Workflow) *WorkflowRepostitory {
	repo.where.Push(&workflow)
	return repo
}

func (repo *WorkflowRepostitory) First() (*Workflow, error) {
	defer func() {
		repo.db.Close()
	}()
	workflow := &Workflow{}
	err := repo.prepareDb().First(workflow).Error
	repo.ResetInstance()
	return workflow, err
}

func (repo *WorkflowRepostitory) Update(workflow Workflow, payload Workflow) (*Workflow, error) {
	defer func() {
		repo.db.Close()
	}()

	err := repo.prepareDb().Model(&workflow).Update(payload).Error
	repo.ResetInstance()
	return &workflow, err
}

func (repo *WorkflowRepostitory) Insert(workflow Workflow) (*Workflow, error) {
	defer func() {
		repo.db.Close()
	}()

	in := &workflow
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *WorkflowRepostitory) Delete(workflow Workflow) (*Workflow, error) {
	defer func() {
		repo.db.Close()
	}()
	in := &workflow
	if in.ID == "" {
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
	return tx
}

func (repo *WorkflowRepostitory) ResetInstance() {
	repo.adapter = nil
}

func NewWorkflowRepository() *WorkflowRepostitory {
	db, _ := db.NewEngine()
	return &WorkflowRepostitory{
		db:    db,
		where: &stack.Stack{},
	}
}
