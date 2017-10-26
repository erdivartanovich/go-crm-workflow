package workflow

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
	return service.Repo.Update(workflow, payload)
}

func (service *WorkflowService) Add(workflow Workflow) (*Workflow, error) {
	return service.Repo.Insert(workflow)
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
