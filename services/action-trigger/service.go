package actiontrigger

import (
	"fmt"

	"github.com/kwri/go-workflow/services/entity"
)

type ActionTriggerService struct {
	Repo *ActionTriggerRepository
}

func (service *ActionTriggerService) Browse(adapter *entity.SearchAdapter) ([]*entity.ActionTrigger, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *ActionTriggerService) Read(actiontrigger entity.ActionTrigger) (*entity.ActionTrigger, error) {
	return service.Repo.Where(actiontrigger).First()
}

func (service *ActionTriggerService) Edit(actiontrigger entity.ActionTrigger, payload entity.ActionTrigger) (*entity.ActionTrigger, error) {
	at, err := service.Read(actiontrigger)
	if err != nil {
		return nil, err
	}
	return service.Repo.Update(*at, payload)
}

func (service *ActionTriggerService) Replace(actiontrigger entity.ActionTrigger, payload entity.ActionTrigger) (*entity.ActionTrigger, error) {
	return service.Repo.Replace(actiontrigger, payload)
}

func (service *ActionTriggerService) Add(actiontrigger entity.ActionTrigger) (*entity.ActionTrigger, error) {
	return service.Repo.Insert(actiontrigger)
}

func (service *ActionTriggerService) BatchAdd(payloads []entity.ActionTrigger) (int, error) {
	var ch chan bool
	ch = make(chan bool)

	go func() {
		for i := range payloads {
			fmt.Println("Lets save", i)
			_, err := service.Add(payloads[i])
			if err != nil {
				ch <- false
				continue
			}
			ch <- true
		}
		fmt.Println("closed")
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

func (service *ActionTriggerService) Delete(actiontrigger entity.ActionTrigger) (*entity.ActionTrigger, error) {
	return service.Repo.Delete(actiontrigger)
}

func (service *ActionTriggerService) Count(adapter *entity.SearchAdapter) (int, error) {
	return service.Repo.SetAdapter(adapter).Count()
}

func NewActionTriggerService() *ActionTriggerService {
	return &ActionTriggerService{
		Repo: NewActionTriggerRepository(),
	}
}
