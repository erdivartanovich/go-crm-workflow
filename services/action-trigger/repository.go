package actiontrigger

import (
	"errors"
	"github.com/golang-collections/collections/stack"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
)

type ActionTriggerRepository struct {
	db 		*gorm.DB
	adapter *SearchAdapter
	where *stack.Stack
}

func (repo *ActionTriggerRepository) SetAdapter(adapter SearchAdapter) *ActionTriggerRepository {
	repo.adapter = &adapter
	return repo
}

func (repo *ActionTriggerRepository) Find() ([]*ActionTrigger, error) {
	ats := &[]*ActionTrigger{}
	err := repo.preparedDb().Find(ats).Error
	repo.ResetInstance()
	return *ats, err
}

func (repo *ActionTriggerRepository) Where(actiontrigger ActionTrigger) *ActionTriggerRepository {
	repo.where.Push(&ActionTrigger)
	return repo
}

func (actiontrigger *ActionTriggerRepository) First() (*ActionTrigger, error) {
	at := &ActionTrigger{}
	err := repo.preparedDb().First(at).Error()
	repo.ResetInstance()
	return at, err
}

func (repo *ActionTriggerRepository) Update(actiontrigger ActionTrigger, payload ActionTrigger) (*ActionTrigger, error) {
	err := repo.preparedDb().Model(&actiontrigger).Update(payload).Error
	repo.ResetInstance()
	return &action, err
}

func (repo *ActionTriggerRepository) Replace(actiontrigger ActionTrigger, payload ActionTrigger) (*ActionTrigger, error) {
	at := &actiontrigger
	db := repo.preparedDb()
	db.First(a)
	at.Name = payload.Name
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

func (repo *ActionTriggerRepository) Insert(actiontrigger ActionTrigger) (*ActionTrigger, error) {
	in := &actiontrigger
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}
	
func (repo *ActionTriggerRepository) Delete(actiontrigger ActionTriggerion) (*ActionTrigger, error) {
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
	err := repo.prepareDb().Model(&ActionTrigger{}).Count(&count).Error
	return count, err
}

func NewActionTriggerRepository() *ActionTriggerRepository {
	db := db.Engine
	return &ActionTriggerRepository{
		db:    db,
		where: &stack.Stack{},
	}
}