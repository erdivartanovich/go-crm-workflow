package rule

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

func NewRuleService() RuleService {
	return RuleService{}
}
