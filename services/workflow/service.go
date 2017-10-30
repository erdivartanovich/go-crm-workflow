package workflow

import "fmt"

type WorkflowService struct {
	Repo *WorkflowRepostitory
}

type SearchAdapter struct {
}

func (service *WorkflowService) Browse(adapter SearchAdapter) ([]*Workflow, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *WorkflowService) Read(workflow Workflow) (*Workflow, error) {
	return service.Repo.Where(workflow).First()
}

func (service *WorkflowService) Edit(workflow Workflow, payload Workflow) (*Workflow, error) {
	wk, err := service.Read(workflow)
	if err != nil {
		return nil, err
	}
	return service.Repo.Update(*wk, payload)
}

func (service *WorkflowService) Replace(workflow Workflow, payload Workflow) (*Workflow, error) {
	return service.Repo.Replace(workflow, payload)
}

func (service *WorkflowService) Add(workflow Workflow) (*Workflow, error) {
	return service.Repo.Insert(workflow)
}

func (service *WorkflowService) BatchAdd(payloads []Workflow) (int, error) {
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

func (service *WorkflowService) Delete(workflow Workflow) (*Workflow, error) {
	return service.Repo.Delete(workflow)
}

func (service *WorkflowService) Count(adapter SearchAdapter) (int, error) {
	return service.Repo.SetAdapter(adapter).Count()
}

func NewWorkflowService() *WorkflowService {
	return &WorkflowService{
		Repo: NewWorkflowRepository(),
	}
}
