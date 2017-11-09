package rule

import "github.com/kwri/go-workflow/services/entity"

type RuleService struct {
	Repo *RuleRepository
}

func (service *RuleService) Browse(adapter *entity.SearchAdapter) ([]*entity.Rule, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *RuleService) Read(rule entity.Rule) (*entity.Rule, error) {
	return service.Repo.Where(rule).First()
}

func (service *RuleService) Edit(rule entity.Rule, payload entity.Rule) (*entity.Rule, error) {
	rl, err := service.Read(rule)

	if err != nil {
		return nil, err
	}
	return service.Repo.Update(*rl, payload)
}

func (service *RuleService) Replace(rule entity.Rule, playload entity.Rule) (*entity.Rule, error) {
	return service.Repo.Replace(rule, playload)
}

func (service *RuleService) Add(rule entity.Rule) (*entity.Rule, error) {
	return service.Repo.Insert(rule)
}

func (service *RuleService) BatchAdd(playloads []entity.Rule) (int, error) {
	var ch chan bool
	ch = make(chan bool)

	go func() {
		for i := range playloads {
			_, err := service.Add(playloads[i])
			if err != nil {
				ch <- false
				continue
			}
			ch <- true
		}
		close(ch)
	}()
	success := 0
	for n := range ch {
		if n == true {
			success++
		}
	}

	return success, nil
}

func (service *RuleService) Delete(rule entity.Rule) (*entity.Rule, error) {
	return service.Repo.Delete(rule)
}

func (service *RuleService) Count(adapter *entity.SearchAdapter) (int, error) {
	return service.Repo.SetAdapter(adapter).Count()
}

func NewRuleService() *RuleService {
	return &RuleService{
		Repo: NewRuleRepository(),
	}
}
