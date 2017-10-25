package rule

import (
	"github.com/kwri/go-workflow/services/action"
)

type RuleService struct {
}

type SearchAdapter struct {
}

var repo = NewRuleRepository()

func (service RuleService) Browse(adapter *SearchAdapter) (*[]Rule, error) {
	return repo.SetAdapter(adapter).Find()
}

func (service RuleService) Read(rule Rule) (*Rule, error) {
	return repo.Where(rule).First()
}

func (service RuleService) Edit(rule Rule) (*Rule, error) {
	return repo.Update(rule)
}

func (service RuleService) Add(rule Rule) (*Rule, error) {
	return repo.Insert(rule)
}

func (service RuleService) Delete(rule Rule) (*Rule, error) {
	return repo.Delete(rule)
}

func (service RuleService) syncActions(rule Rule, actions ...action.Action) (*Rule, error) {
	return repo.syncActions(rule, actions...)
}

func NewRuleService() RuleService {
	return RuleService{}
}
