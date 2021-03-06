package action

import (
	"errors"

	"github.com/golang-collections/collections/stack"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/services/entity"
)

type ActionRepository struct {
	db      *gorm.DB
	adapter *entity.SearchAdapter
	where   *stack.Stack
}

func (repo *ActionRepository) SetAdapter(adapter entity.SearchAdapter) *ActionRepository {
	repo.adapter = &adapter
	return repo
}

func (repo *ActionRepository) Find() ([]*entity.Action, error) {
	actions := &[]*entity.Action{}
	err := repo.prepareDb().Find(actions).Error
	repo.ResetInstance()
	return *actions, err
}

func (repo *ActionRepository) Where(action entity.Action) *ActionRepository {
	repo.where.Push(&action)
	return repo
}

func (repo *ActionRepository) First() (*entity.Action, error) {
	action := &entity.Action{}
	err := repo.prepareDb().First(action).Error
	repo.ResetInstance()
	return action, err
}

func (repo *ActionRepository) Update(action entity.Action, payload entity.Action) (*entity.Action, error) {
	payload.ID = action.ID
	err := repo.prepareDb().Model(&action).Update(payload).Error
	repo.ResetInstance()
	return &action, err
}

func (repo *ActionRepository) Replace(action entity.Action, payload entity.Action) (*entity.Action, error) {
	a := &action
	db := repo.prepareDb()
	db.First(a)
	a.Name = payload.Name
	a.ActionType = payload.ActionType
	a.TargetClass = payload.TargetClass
	a.TargetField = payload.TargetField
	a.Value = payload.Value
	err := db.Save(a).Error
	repo.ResetInstance()
	return a, err
}

func (repo *ActionRepository) Insert(action entity.Action) (*entity.Action, error) {
	in := &action
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *ActionRepository) Delete(action entity.Action) (*entity.Action, error) {
	in := &action
	if len(in.ID) == 0 {
		return nil, errors.New("You need to set ID of deleted action")
	}
	err := repo.prepareDb().Delete(&in).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Unscoped().Find(&in).Error
	repo.ResetInstance()
	return in, err
}
func (repo *ActionRepository) prepareDb() *gorm.DB {
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

func (repo *ActionRepository) ResetInstance() {
	repo.adapter = nil
}

func (repo *ActionRepository) Count() (int, error) {
	count := 0
	err := repo.prepareDb().Model(&entity.Action{}).Count(&count).Error
	return count, err
}

func NewActionRepository() *ActionRepository {
	db := db.Engine
	return &ActionRepository{
		db:    db,
		where: &stack.Stack{},
	}
}
