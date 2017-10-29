package action

import (
	"fmt"
)

type ActionService struct {
	Repo *ActionRepository
}

type SearchAdapter struct {
}

func (service *ActionService) Browse(adapter SearchAdapter) ([]*Action, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *ActionService) Read(action Action) (*Action, error) {
	return service.Repo.Where(action).First()
}

func (service *ActionService) Edit(action Action, payload Action) (*Action, error) {
	wk, err := service.Read(action)
	if err != nil {
		return nil, err
	}
	return service.Repo.Update(*wk, payload)
}

func (service *ActionService) Replace(action Action, payload Action) (*Action, error) {
	return service.Repo.Replace(action, payload)
}

func (service *ActionService) Add(action Action) (*Action, error) {
	return service.Repo.Insert(action)
}

func (service *ActionService) BatchAdd(payloads []Action) (int, error) {
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

func (service *ActionService) Delete(action Action) (*Action, error) {
	return service.Repo.Delete(action)
}

func (service *ActionService) Count(adapter SearchAdapter) (int, error) {
	return service.Repo.SetAdapter(adapter).Count()
}

func NewActionService() *ActionService {
	return &ActionService{
		Repo:NewActionRepository(),
	}
}