package v1

import (
	"net/http"

	"github.com/kwri/go-workflow/services/workflow"
	api "github.com/kwri/go-workflow/vndapi"
)

type workflowCtrl struct {
	service *workflow.WorkflowService
}

func newWorkflowCtrl() *workflowCtrl {
	return &workflowCtrl{
		service: workflow.NewWorkflowService(),
	}
}

func (ctrl *workflowCtrl) Browse(r *http.Request) (api.Responder, error) {
	service := workflow.NewWorkflowService()
	adapter := workflow.SearchAdapter{}
	workflows, err := service.Browse(adapter)
	return &api.ApiResponder{
		Data:     workflows,
		Hostname: "https://localhost:8001",
	}, err
}

func (ctrl *workflowCtrl) Read(id string, r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *workflowCtrl) Replace(id string, r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *workflowCtrl) Edit(id string, r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *workflowCtrl) Add(r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *workflowCtrl) Delete(id string, r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *workflowCtrl) BatchAdd(r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *workflowCtrl) BatchEdit(r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *workflowCtrl) Destroy(r *http.Request) (api.Responder, error) {
	return nil, nil
}
