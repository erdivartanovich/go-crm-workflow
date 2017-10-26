package rule

import (
	"github.com/kwri/go-workflow/services/action"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
)

type RuleRepostitory struct {
	db      *gorm.DB
	adapter *SearchAdapter
}

func (repo RuleRepostitory) SetAdapter(adapter *SearchAdapter) RuleRepostitory {
	repo.adapter = adapter
	return repo
}

func (repo RuleRepostitory) Find() (*[]Rule, error) {
	defer func() {
		repo.db.Close()
	}()
	rules := &[]Rule{}
	err := repo.db.Find(rules).Error
	return rules, err
}

func (repo RuleRepostitory) Where(rule Rule) RuleRepostitory {
	repo.db.Where(rule)
	return repo
}

func (repo RuleRepostitory) First() (*Rule, error) {
	defer func() {
		repo.db.Close()
	}()
	rule := &Rule{}
	err := repo.db.First(rule).Error
	return rule, err
}

func (repo RuleRepostitory) Update(rule Rule) (*Rule, error) {
	defer func() {
		repo.db.Close()
	}()
	in := &Rule{}
	err := repo.db.Save(in).Error
	return in, err
}

func (repo RuleRepostitory) Insert(rule Rule) (*Rule, error) {
	defer func() {
		repo.db.Close()
	}()

	if !repo.db.NewRecord(rule) {
		return nil, errors.New(fmt.Sprintf(
			"User with id %s is exists on database",
			rule.ID,
		))
	}
	in := &rule
	err := repo.db.Create(in).Error
	return in, err
}

func (repo RuleRepostitory) Delete(rule Rule) (*Rule, error) {
	defer func() {
		repo.db.Close()
	}()
	in := &rule
	err := repo.db.Delete(in).Error
	return in, err
}

func (repo RuleRepostitory) syncActions(rule Rule, actions ...action.Action) (*Rule, error) {
	defer func() {
		repo.db.Close()
	}()
	// get the action ids
	var ids [][]byte
	for _, action := range(actions) {
		ids = append(ids, action.ID)
	}
	repo.db.Model(&rule).Related(&actions, "actions")
	return in, err
}

func NewRuleRepository() *RuleRepostitory {
	return &RuleRepostitory{
		db: db.Engine,
	}
}
