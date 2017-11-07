package workflow

import (
	"github.com/kwri/go-workflow/services/entity"
)

type WorkflowService struct {
	Repo *WorkflowRepostitory
}

func (service *WorkflowService) Browse(adapter *entity.SearchAdapter) ([]*entity.Workflow, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *WorkflowService) Read(workflow entity.Workflow) (*entity.Workflow, error) {
	return service.Repo.Where(workflow).First()
}

func (service *WorkflowService) Edit(workflow entity.Workflow, payload entity.Workflow) (*entity.Workflow, error) {
	wk, err := service.Read(workflow)

	if err != nil {
		return nil, err
	}
	return service.Repo.Update(*wk, payload)
}

func (service *WorkflowService) Replace(workflow entity.Workflow, payload entity.Workflow) (*entity.Workflow, error) {
	return service.Repo.Replace(workflow, payload)
}

func (service *WorkflowService) Add(workflow entity.Workflow) (*entity.Workflow, error) {
	return service.Repo.Insert(workflow)
}

func (service *WorkflowService) BatchAdd(payloads []entity.Workflow) (int, error) {
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

func (service *WorkflowService) Delete(workflow entity.Workflow) (*entity.Workflow, error) {
	return service.Repo.Delete(workflow)
}

func (service *WorkflowService) Count(adapter *entity.SearchAdapter) (int, error) {
	return service.Repo.SetAdapter(adapter).Count()
}

func NewWorkflowService() *WorkflowService {
	return &WorkflowService{
		Repo: NewWorkflowRepository(),
	}
}
