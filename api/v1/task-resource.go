package v1

import (
	"io/ioutil"
	"net/http"

	paginator "github.com/kwri/go-workflow/gorm-paginator"
	"github.com/kwri/go-workflow/services/entity"
	"github.com/kwri/go-workflow/services/task"
	api "github.com/kwri/go-workflow/vndapi"
	"github.com/manyminds/api2go/jsonapi"
)

type taskCtrl struct {
	service *task.TaskService
}

func newTaskCtrl() *taskCtrl {
	return &taskCtrl{
		service: task.NewTaskService(),
	}
}

func (ctrl *taskCtrl) Browse(r *http.Request) (api.Responder, error) {
	service := ctrl.service
	adapter := entity.ToSearchAdapter(r.URL.Query())
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

	var tasks []*entity.Task

	if total > 0 {
		tasks, err = service.Browse(adapter)
	}

	paginator := paginator.NewLengthAwareOffsetPaginator(tasks, total, limit, offset, options)

	respond := &api.ApiResponder{
		Data: paginator,
		Code: 200,
	}

	return respond, err
}

func (ctrl *taskCtrl) Read(id string, r *http.Request) (api.Responder, error) {
	service := ctrl.service
	payload := &entity.Task{}
	payload.SetID(id)
	task, err := service.Read(*payload)

	return &api.ApiResponder{
		Data: task,
		Code: 200,
	}, err
}

func (ctrl *taskCtrl) Replace(id string, r *http.Request) (api.Responder, error) {
	wk := entity.Task{}
	wk.SetID(id)
	payload := entity.Task{}
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

func (ctrl *taskCtrl) Edit(id string, r *http.Request) (api.Responder, error) {
	wk := entity.Task{}
	wk.SetID(id)
	payload := entity.Task{}
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

func (ctrl *taskCtrl) Add(r *http.Request) (api.Responder, error) {
	payload := entity.Task{}
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

	task, err := ctrl.service.Add(payload)

	return &api.ApiResponder{
		Data: task,
		Code: 200,
	}, err
}

func (ctrl *taskCtrl) Delete(id string, r *http.Request) (api.Responder, error) {
	wk := entity.Task{}
	wk.SetID(id)

	_, err := ctrl.service.Delete(wk)

	return &api.ApiResponder{
		Data: nil,
		Code: 204,
	}, err
}

func (ctrl *taskCtrl) BatchAdd(r *http.Request) (api.Responder, error) {
	var payloads []entity.Task
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

func (ctrl *taskCtrl) BatchEdit(r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *taskCtrl) Destroy(r *http.Request) (api.Responder, error) {
	return nil, nil
}
