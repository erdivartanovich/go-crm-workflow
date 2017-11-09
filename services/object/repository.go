package object

import (
	"errors"

	"github.com/golang-collections/collections/stack"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/services/entity"
)

var IncludeMap = map[string]string{
	"actions": "Actions",
}

type WorkflowObjectRepository struct {
	db      *gorm.DB
	adapter *entity.SearchAdapter
	where   *stack.Stack
}

func (repo *WorkflowObjectRepository) SetAdapter(adapter *entity.SearchAdapter) *WorkflowObjectRepository {
	repo.adapter = adapter
	return repo
}

func (repo *WorkflowObjectRepository) Find() ([]*entity.WorkflowObject, error) {
	workflows := &[]*entity.WorkflowObject{}
	err := repo.prepareDb().Find(workflows).Error
	repo.ResetInstance()
	return *workflows, err
}

func (repo *WorkflowObjectRepository) Where(workflow entity.WorkflowObject) *WorkflowObjectRepository {
	repo.where.Push(&workflow)
	return repo
}

func (repo *WorkflowObjectRepository) First() (*entity.WorkflowObject, error) {

	workflow := &entity.WorkflowObject{}
	err := repo.prepareDb().First(workflow).Error
	repo.ResetInstance()
	return workflow, err
}

func (repo *WorkflowObjectRepository) Update(workflow entity.WorkflowObject, payload entity.WorkflowObject) (*entity.WorkflowObject, error) {

	err := repo.prepareDb().Model(&workflow).Update(payload).Error
	repo.ResetInstance()
	return &workflow, err
}

func (repo *WorkflowObjectRepository) Replace(workflow entity.WorkflowObject, payload entity.WorkflowObject) (*entity.WorkflowObject, error) {
	wk := &workflow
	db := repo.prepareDb()
	db.First(wk)
	wk.ObjectClass = payload.ObjectClass
	wk.ObjectType = payload.ObjectType
	err := db.Save(wk).Error
	repo.ResetInstance()
	return wk, err
}

func (repo *WorkflowObjectRepository) Insert(workflow entity.WorkflowObject) (*entity.WorkflowObject, error) {

	in := &workflow
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *WorkflowObjectRepository) Delete(workflow entity.WorkflowObject) (*entity.WorkflowObject, error) {

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

func (repo *WorkflowObjectRepository) prepareDb() *gorm.DB {
	count := repo.where.Len()
	tx := repo.db
	for i := 0; i < count; i++ {
		tx = tx.Where(repo.where.Pop())
	}

	tx = repo.applySearchAdapter(tx)

	return tx
}

func (repo *WorkflowObjectRepository) applySearchAdapter(tx *gorm.DB) *gorm.DB {
	if repo.adapter != nil {
		tx = repo.applyInclude(tx)
		tx = repo.applyFilters(tx)
		tx = repo.applySorter(tx)
		tx = repo.applyPager(tx)
	}
	return tx
}

func (repo *WorkflowObjectRepository) applyInclude(tx *gorm.DB) *gorm.DB {

	if repo.adapter.Include != nil && len(repo.adapter.Include) > 0 {

		for _, resource := range repo.adapter.Include {

			if val, ok := IncludeMap[resource]; ok {
				tx = tx.Preload(val)
			}
		}
	}

	return tx
}

func (repo *WorkflowObjectRepository) applyPager(tx *gorm.DB) *gorm.DB {
	limit := 10
	offset := 0
	if repo.adapter.Page != nil {
		limit = repo.adapter.Page.Limit
		offset = repo.adapter.Page.Offset
	}
	tx = tx.Limit(limit).Offset(offset)
	return tx
}

func (repo *WorkflowObjectRepository) applyFilters(tx *gorm.DB) *gorm.DB {
	if repo.adapter.Filters != nil && len(repo.adapter.Filters) > 0 {

	}
	return tx
}

func (repo *WorkflowObjectRepository) applySorter(tx *gorm.DB) *gorm.DB {
	if repo.adapter.Sort != "" {
		tx = tx.Order(string(repo.adapter.Sort))
	}
	return tx
}

func (repo *WorkflowObjectRepository) ResetInstance() {
	repo.adapter = nil
}

func (repo *WorkflowObjectRepository) Count() (int, error) {
	count := 0
	tx := repo.prepareDb().Limit(-1).Offset(-1)
	err := tx.Model(&entity.WorkflowObject{}).Count(&count).Error
	return count, err
}

func NewWorkflowRepository() *WorkflowObjectRepository {
	db := db.Engine
	return &WorkflowObjectRepository{
		db:    db,
		where: &stack.Stack{},
	}
}
