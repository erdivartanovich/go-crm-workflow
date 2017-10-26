package v1

import (
	"fmt"
	"net/http"
	"strconv"

	paginator "github.com/kwri/go-workflow/gorm-paginator"
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
	service := ctrl.service
	adapter := workflow.SearchAdapter{}
	workflows, err := service.Browse(adapter)
	total, err := service.Count(adapter)
	if err != nil {
		total = 0
	}
	qlimit := r.URL.Query().Get("page[limit]")

	limit := 10
	if qlimit != "" {

		val, e := strconv.Atoi(qlimit)

		if e == nil {
			limit = val
		}
	}
	qoffset := r.URL.Query().Get("page[offset]")
	offset := 0
	if qoffset != "" {
		val, e := strconv.Atoi(qoffset)
		if e == nil {
			offset = val
		}
	}
	options := &paginator.Options{}
	paginator := paginator.NewLengthAwareOffsetPaginator(workflows, total, limit, offset, options)
	respond := &api.ApiResponder{
		Data:     paginator,
		Hostname: "https://localhost:8001",
	}

	return respond, err
}

func (ctrl *workflowCtrl) Read(id string, r *http.Request) (api.Responder, error) {

	service := ctrl.service
	payload := &workflow.Workflow{}
	payload.UnmarshalUUIDString(id)
	workflow, err := service.Read(*payload)
	fmt.Println(err)
	return &api.ApiResponder{
		Data:     workflow,
		Hostname: "https://localhost:8001",
	}, err
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
