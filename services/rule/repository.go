package rule

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

type RuleRepository struct {
	db      *gorm.DB
	adapter *entity.SearchAdapter
	where   *stack.Stack
}

func (repo *RuleRepository) SetAdapter(adapter *entity.SearchAdapter) *RuleRepository {
	repo.adapter = adapter
	return repo
}

func (repo *RuleRepository) Find() ([]*entity.Rule, error) {
	rules := &[]*entity.Rule{}
	err := repo.prepareDb().Find(rules).Error
	repo.ResetInstance()
	return *rules, err
}

func (repo *RuleRepository) Where(rule entity.Rule) *RuleRepository {
	repo.where.Push(&rule)
	return repo
}

func (repo *RuleRepository) First() (*entity.Rule, error) {
	rule := &entity.Rule{}
	err := repo.prepareDb().First(rule).Error
	repo.ResetInstance()
	return rule, err
}

func (repo *RuleRepository) Update(rule entity.Rule, payload entity.Rule) (*entity.Rule, error) {
	payload.ID = rule.ID
	err := repo.prepareDb().Model(&rule).Update(payload).Error
	repo.ResetInstance()
	return &rule, err
}

func (repo *RuleRepository) Replace(rule entity.Rule, payload entity.Rule) (*entity.Rule, error) {
	rl := &rule
	db := repo.prepareDb()
	db.First(rl)
	rl.UserID = payload.UserID
	rl.Name = payload.Name
	rl.RuleType = payload.RuleType
	rl.FieldName = payload.FieldName
	rl.Operator = payload.Operator
	rl.Value = payload.Value
	rl.Priority = payload.Priority
	rl.Actions = payload.Actions
	err := db.Save(rl).Error
	repo.ResetInstance()
	return rl, err
}

func (repo *RuleRepository) Insert(rule entity.Rule) (*entity.Rule, error) {
	in := &rule
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *RuleRepository) Delete(rule entity.Rule) (*entity.Rule, error) {
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

	if repo.adapter != nil {
		tx = repo.adapter.ApplySearchAdapter(tx)
	}

	return tx
}

func (repo *RuleRepository) ResetInstance() {
	repo.adapter = nil
}

func (repo *RuleRepository) Count() (int, error) {
	count := 0
	tx := repo.prepareDb().Limit(-1).Offset(-1)
	err := tx.Model(&entity.Rule{}).Count(&count).Error
	return count, err
}

func NewRuleRepository() *RuleRepository {
	db := db.Engine
	return &RuleRepository{
		db:    db,
		where: &stack.Stack{},
	}
}
