package action

import (
	"errors"
	"github.com/golang-collections/collections/stack"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
)
type ActionRepository struct {
	db *gorm.DB
	adapter *SearchAdapter
	where *stack.Stack
}

func (repo *ActionRepository) SetAdapter(adapter SearchAdapter) *ActionRepository {
	repo.adapter = &adapter
	return repo
}

func (repo *ActionRepository) Find() ([]*Action, error) {
	actions := &[]*Action{}
	err := repo.prepareDb().Find(action).Error
	repo.ResetInstance()
	return *actions, err
}

func (repo *ActionRepository) Where(action Action) *ActionRepository {
	repo.where.Push(&action)
	return repo
}

func (repo *ActionRepository) First() (*Action, error) {
	action := &Action{}
	err := repo.prepareDb().First(action).Error
	repo.ResetInstance()
	return action, err
}

func (repo *ActionRepository) Update(action Action, payload Action) (*Action, error) {
	err := repo.prepareDb().Model(&action).Update(payload).Error
	repo.ResetInstance()
	return &action, err
}

func (repo *ActionRepostitory) Insert(action Action) (*Action, error) {
	in := &action
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}
	
func (repo *ActionRepostitory) Delete(action Action) (*Action, error) {
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
func (repo *ActionRepostitory) prepareDb() *gorm.DB {
	count := repo.where.Len()
	tx := repo.db
	for i := 0; i < count; i++ {
		tx = tx.Where(repo.where.Pop())
	}
	return tx
}

func (repo *ActionRepostitory) ResetInstance() {
	repo.adapter = nil
}

func (repo *ActionRepostitory) Count() (int, error) {
	count := 0
	err := repo.prepareDb().Model(&Action{}).Count(&count).Error
	return count, err
}

func NewActionRepository() *ActionRepostitory {
	db := db.Engine
	return &ActionRepostitory{
		db:    db,
		where: &stack.Stack{},
	}
}
	