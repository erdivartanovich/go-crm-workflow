package v1

import (
	"io/ioutil"
	"net/http"

	paginator "github.com/kwri/go-workflow/gorm-paginator"
	"github.com/kwri/go-workflow/services/entity"
	"github.com/kwri/go-workflow/services/object"

	api "github.com/kwri/go-workflow/vndapi"
	"github.com/manyminds/api2go/jsonapi"
)

type workflowObjectCtrl struct {
	service *object.WorkflowObjectService
}

func newWorkflowObjectCtrl() *workflowObjectCtrl {
	return &workflowObjectCtrl{
		service: object.NewWorkflowObjectService(),
	}
}

func (ctrl *workflowObjectCtrl) Browse(r *http.Request) (api.Responder, error) {
	service := ctrl.service
	adapter := &entity.SearchAdapter{}
	adapter.FromURLValues(r.URL.Query())
	total, err := service.Count(adapter)

	if err != nil {
		total = 0
	}

	limit := adapter.Page.Limit
	offset := adapter.Page.Offset
	options := &paginator.Options{
		QueryParameter: r.URL.Query(),
		Path:           r.URL.Path,
	}

	var workflows []*entity.WorkflowObject

	if total > 0 {
		workflows, err = service.Browse(adapter)
	}

	paginator := paginator.NewLengthAwareOffsetPaginator(workflows, total, limit, offset, options)

	respond := &api.ApiResponder{
		Data: paginator,
		Code: 200,
	}

	return respond, err
}

func (ctrl *workflowObjectCtrl) Read(id string, r *http.Request) (api.Responder, error) {
	service := ctrl.service
	payload := &entity.WorkflowObject{}
	payload.SetID(id)
	workflow, err := service.Read(*payload)

	return &api.ApiResponder{
		Data: workflow,
		Code: 200,
	}, err
}

func (ctrl *workflowObjectCtrl) Replace(id string, r *http.Request) (api.Responder, error) {
	wk := entity.WorkflowObject{}
	wk.SetID(id)
	payload := entity.WorkflowObject{}
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

func (ctrl *workflowObjectCtrl) Edit(id string, r *http.Request) (api.Responder, error) {
	wk := entity.WorkflowObject{}
	wk.SetID(id)
	payload := entity.WorkflowObject{}
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

	updated, err := ctrl.service.Edit(wk, payload)

	return &api.ApiResponder{
		Data: updated,
		Code: 200,
	}, err
}

func (ctrl *workflowObjectCtrl) Add(r *http.Request) (api.Responder, error) {
	payload := entity.WorkflowObject{}
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

	workflow, err := ctrl.service.Add(payload)

	return &api.ApiResponder{
		Data: workflow,
		Code: 200,
	}, err
}

func (ctrl *workflowObjectCtrl) Delete(id string, r *http.Request) (api.Responder, error) {
	wk := entity.WorkflowObject{}
	wk.SetID(id)

	_, err := ctrl.service.Delete(wk)

	return &api.ApiResponder{
		Data: nil,
		Code: 204,
	}, err
}

func (ctrl *workflowObjectCtrl) BatchAdd(r *http.Request) (api.Responder, error) {
	var payloads []entity.WorkflowObject
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &api.ApiResponder{
			Data: nil,
			Code: 422,
		}, err
	}

	err = jsonapi.Unmarshal(body, &payloads)
	if err != nil {

		return &api.ApiResponder{
			Data: nil,
			Code: 422,
		}, err
	}

	success, err := ctrl.service.BatchAdd(payloads)

	return &api.ApiResponder{
		Meta: map[string]interface{}{
			"saved_count": success,
		},
		Data: nil,
		Code: 200,
	}, err
}

func (ctrl *workflowObjectCtrl) BatchEdit(r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *workflowObjectCtrl) Destroy(r *http.Request) (api.Responder, error) {
	return nil, nil
}
