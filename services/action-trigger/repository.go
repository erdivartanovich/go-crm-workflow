package actiontrigger

import (
	"errors"

	"github.com/golang-collections/collections/stack"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/services/entity"
)

type ActionTriggerRepository struct {
	db      *gorm.DB
	adapter *entity.SearchAdapter
	where   *stack.Stack
}

func (repo *ActionTriggerRepository) SetAdapter(adapter *entity.SearchAdapter) *ActionTriggerRepository {
	repo.adapter = adapter
	return repo
}

func (repo *ActionTriggerRepository) Find() ([]*entity.ActionTrigger, error) {
	ats := &[]*entity.ActionTrigger{}
	err := repo.prepareDb().Find(ats).Error
	repo.ResetInstance()
	return *ats, err
}

func (repo *ActionTriggerRepository) Where(actiontrigger entity.ActionTrigger) *ActionTriggerRepository {
	repo.where.Push(&actiontrigger)
	return repo
}

func (repo *ActionTriggerRepository) First() (*entity.ActionTrigger, error) {
	at := &entity.ActionTrigger{}
	err := repo.prepareDb().First(at).Error
	repo.ResetInstance()
	return at, err
}

func (repo *ActionTriggerRepository) Update(actiontrigger entity.ActionTrigger, payload entity.ActionTrigger) (*entity.ActionTrigger, error) {
	payload.ID = actiontrigger.ID
	err := repo.prepareDb().Model(&actiontrigger).Update(payload).Error
	repo.ResetInstance()
	return &actiontrigger, err
}

func (repo *ActionTriggerRepository) Replace(actiontrigger entity.ActionTrigger, payload entity.ActionTrigger) (*entity.ActionTrigger, error) {
	at := &actiontrigger
	db := repo.prepareDb()
	db.First(at)
	at.TargetField = payload.TargetField
	at.Hour = payload.Hour
	at.Min = payload.Min
	at.Month = payload.Month
	at.DayPerMonth = payload.DayPerMonth
	at.DayPerWeek = payload.DayPerWeek
	at.RunnableOnce = payload.RunnableOnce
	err := db.Save(at).Error
	repo.ResetInstance()
	return at, err
}

func (repo *ActionTriggerRepository) Insert(actiontrigger entity.ActionTrigger) (*entity.ActionTrigger, error) {
	in := &actiontrigger
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *ActionTriggerRepository) Delete(actiontrigger entity.ActionTrigger) (*entity.ActionTrigger, error) {
	in := &actiontrigger
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

func (repo *ActionTriggerRepository) prepareDb() *gorm.DB {
	count := repo.where.Len()
	tx := repo.db
	for i := 0; i < count; i++ {
		tx = tx.Where(repo.where.Pop())
	}
	return tx
}

func (repo *ActionTriggerRepository) ResetInstance() {
	repo.adapter = nil
}

func (repo *ActionTriggerRepository) Count() (int, error) {
	count := 0
	err := repo.prepareDb().Model(&entity.ActionTrigger{}).Count(&count).Error
	return count, err
}

func NewActionTriggerRepository() *ActionTriggerRepository {
	db := db.Engine
	return &ActionTriggerRepository{
		db:    db,
		where: &stack.Stack{},
	}
}
