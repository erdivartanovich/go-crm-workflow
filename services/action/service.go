package action

import (
	"github.com/kwri/go-workflow/services/entity"
)

type ActionService struct {
	Repo *ActionRepository
}

func (service *ActionService) Browse(adapter entity.SearchAdapter) ([]*entity.Action, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *ActionService) Read(action entity.Action) (*entity.Action, error) {
	return service.Repo.Where(action).First()
}

func (service *ActionService) Edit(action entity.Action, payload entity.Action) (*entity.Action, error) {
	a, err := service.Read(action)
	if err != nil {
		return nil, err
	}
	return service.Repo.Update(*a, payload)
}

func (service *ActionService) Replace(action entity.Action, payload entity.Action) (*entity.Action, error) {
	return service.Repo.Replace(action, payload)
}

func (service *ActionService) Add(action entity.Action) (*entity.Action, error) {
	return service.Repo.Insert(action)
}

func (service *ActionService) BatchAdd(payloads []entity.Action) (int, error) {
	var ch chan bool
	ch = make(chan bool)

	go func() {
		for i := range payloads {

			_, err := service.Add(payloads[i])
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

func (service *ActionService) Delete(action entity.Action) (*entity.Action, error) {
	return service.Repo.Delete(action)
}

func (service *ActionService) Count(adapter entity.SearchAdapter) (int, error) {
	return service.Repo.SetAdapter(adapter).Count()
}

func NewActionService() *ActionService {
	return &ActionService{
		Repo: NewActionRepository(),
	}
}
