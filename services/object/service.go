package object

import (
	"github.com/kwri/go-workflow/services/entity"
)

type WorkflowObjectService struct {
	Repo *WorkflowObjectRepository
}

func (service *WorkflowObjectService) Browse(adapter *entity.SearchAdapter) ([]*entity.WorkflowObject, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *WorkflowObjectService) Read(workflow entity.WorkflowObject) (*entity.WorkflowObject, error) {
	return service.Repo.Where(workflow).First()
}

func (service *WorkflowObjectService) Edit(workflow entity.WorkflowObject, payload entity.WorkflowObject) (*entity.WorkflowObject, error) {
	wk, err := service.Read(workflow)

	if err != nil {
		return nil, err
	}
	return service.Repo.Update(*wk, payload)
}

func (service *WorkflowObjectService) Replace(workflow entity.WorkflowObject, payload entity.WorkflowObject) (*entity.WorkflowObject, error) {
	return service.Repo.Replace(workflow, payload)
}

func (service *WorkflowObjectService) Add(workflow entity.WorkflowObject) (*entity.WorkflowObject, error) {
	return service.Repo.Insert(workflow)
}

func (service *WorkflowObjectService) BatchAdd(payloads []entity.WorkflowObject) (int, error) {
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

func (service *WorkflowObjectService) Delete(workflow entity.WorkflowObject) (*entity.WorkflowObject, error) {
	return service.Repo.Delete(workflow)
}

func (service *WorkflowObjectService) Count(adapter *entity.SearchAdapter) (int, error) {
	return service.Repo.SetAdapter(adapter).Count()
}

func NewWorkflowObjectService() *WorkflowObjectService {
	return &WorkflowObjectService{
		Repo: NewWorkflowRepository(),
	}
}
