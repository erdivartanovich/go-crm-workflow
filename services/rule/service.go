package rule

type RuleService struct {
	Repo *RuleRepository
}

type SearchAdapter struct {
}

// var repo = NewRuleRepository()

func (service *RuleService) Browse(adapter SearchAdapter) ([]*Rule, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *RuleService) Read(rule Rule) (*Rule, error) {
	return service.Repo.Where(rule).First()
}

func (service *RuleService) Edit(rule Rule, payload Rule) (*Rule, error) {
	return service.Repo.Update(rule, payload)
}

func (service *RuleService) Add(rule Rule) (*Rule, error) {
	return service.Repo.Insert(rule)
}

func (service *RuleService) Delete(rule Rule) (*Rule, error) {
	return service.Repo.Delete(rule)
}

func NewRuleService() *RuleService {
	return &RuleService{
		Repo: NewRuleRepository(),
	}
}
