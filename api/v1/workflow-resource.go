package v1

import (
	"io/ioutil"
	"net/http"
	"strconv"

	paginator "github.com/kwri/go-workflow/gorm-paginator"
	"github.com/kwri/go-workflow/services/workflow"
	api "github.com/kwri/go-workflow/vndapi"
	"github.com/manyminds/api2go/jsonapi"
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
	options := &paginator.Options{
		QueryParameter: r.URL.Query(),
		Path:           r.URL.Path,
	}
	paginator := paginator.NewLengthAwareOffsetPaginator(workflows, total, limit, offset, options)
	respond := &api.ApiResponder{
		Data: paginator,
		Code: 200,
	}

	return respond, err
}

func (ctrl *workflowCtrl) Read(id string, r *http.Request) (api.Responder, error) {
	service := ctrl.service
	payload := &workflow.Workflow{}
	payload.UnmarshalUUIDString(id)
	workflow, err := service.Read(*payload)

	return &api.ApiResponder{
		Data: workflow,
		Code: 200,
	}, err
}

func (ctrl *workflowCtrl) Replace(id string, r *http.Request) (api.Responder, error) {
	wk := workflow.Workflow{}
	wk.UnmarshalUUIDString(id)
	payload := workflow.Workflow{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &api.ApiResponder{
			Data: nil,
			Code: 422,
		}, err
	}

	err = jsonapi.Unmarshal(body, &payload)
	if err != nil {
		return &api.ApiResponder{
			Data: nil,
			Code: 422,
		}, err
	}

	updated, err := ctrl.service.Replace(wk, payload)

	return &api.ApiResponder{
		Data: updated,
		Code: 200,
	}, err
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
