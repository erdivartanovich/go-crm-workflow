package rule

import (
	"errors"
	"github.com/golang-collections/collections/stack"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
)

type RuleRepository struct {
	db      *gorm.DB
	adapter *SearchAdapter
	where   *stack.Stack
}

func (repo *RuleRepository) SetAdapter(adapter SearchAdapter) *RuleRepository {
	repo.adapter = &adapter
	return repo
}

func (repo *RuleRepository) Find() ([]*Rule, error) {
	rules := &[]*Rule{}
	err := repo.prepareDb().Find(rules).Error
	repo.ResetInstance()
	return *rules, err
}

func (repo *RuleRepository) Where(rule Rule) *RuleRepository {
	repo.where.Push(&rule)
	return repo
}

func (repo *RuleRepository) First() (*Rule, error) {
	rule := &Rule{}
	err := repo.prepareDb().First(rule).Error
	repo.ResetInstance()
	return rule, err
}

func (repo *RuleRepository) Update(rule Rule, payload Rule) (*Rule, error) {
	err := repo.prepareDb().Model(&rule).Update(payload).Error
	repo.ResetInstance()
	return &rule, err
}

func (repo *RuleRepository) Insert(rule Rule) (*Rule, error) {
	in := &rule
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *RuleRepository) Delete(rule Rule) (*Rule, error) {
	in := &rule
	if len(in.ID) == 0 {
		return nil, errors.New("You need to set ID of deleted rule")
	}
	err := repo.prepareDb().Delete(&in).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Unscoped().Find(&in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *RuleRepository) prepareDb() *gorm.DB {
	count := repo.where.Len()
	tx := repo.db
	for i := 0; i < count; i++ {
		tx = tx.Where(repo.where.Pop())
	}
	return tx
}

func (repo *RuleRepository) ResetInstance() {
	repo.adapter = nil
}

func NewRuleRepository() *RuleRepository {
	db := db.Engine
	return &RuleRepository{
		db:    db,
		where: &stack.Stack{},
	}
}
