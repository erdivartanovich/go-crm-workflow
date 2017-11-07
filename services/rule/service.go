package rule

import "github.com/kwri/go-workflow/services/entity"

type RuleService struct {
	Repo *RuleRepository
}

type entity.SearchAdapter struct {
}

// var repo = NewRuleRepository()

func (service *RuleService) Browse(adapter entity.SearchAdapter) ([]*entity.Rule, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *RuleService) Read(rule entity.Rule) (*entity.Rule, error) {
	return service.Repo.Where(rule).First()
}

func (service *RuleService) Edit(rule entity.Rule, payload entity.Rule) (*entity.Rule, error) {
	return service.Repo.Update(rule, payload)
}

func (service *RuleService) Add(rule entity.Rule) (*entity.Rule, error) {
	return service.Repo.Insert(rule)
}

func (service *RuleService) Delete(rule entity.Rule) (*entity.Rule, error) {
	return service.Repo.Delete(rule)
}

func NewRuleService() *RuleService {
	return &RuleService{
		Repo: NewRuleRepository(),
	}
}
