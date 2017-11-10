package v1

import (
	"io/ioutil"
	"net/http"

	paginator "github.com/kwri/go-workflow/gorm-paginator"
	"github.com/kwri/go-workflow/services/action"
	"github.com/kwri/go-workflow/services/entity"
	api "github.com/kwri/go-workflow/vndapi"
	"github.com/manyminds/api2go/jsonapi"
)

type actionCtrl struct {
	service *action.ActionService
}

func newActionCtrl() *actionCtrl {
	return &actionCtrl{
		service: action.NewActionService(),
	}
}

func (ctrl *actionCtrl) Browse(r *http.Request) (api.Responder, error) {
	service := ctrl.service
	adapter := entity.SearchAdapter{}
	adapter.FromURLValues(r.URL.Query())
	total, err := service.Count(adapter)
	
	if err != nil {
		total = 0
	}
	
	limit := adapter.Page.Limit
	offset := adapter.Page.Offset
	options := &paginator.Options{
		QueryParameter: r.URL.Query(),
		Path: 			r.URL.Path,
	}

	var actions []*entity.Action
	
	if total > 0 {
		actions, err = service.Browse(adapter)
	}
	
	paginator := paginator.NewLengthAwareOffsetPaginator(actions, total, limit, offset, options)
	
	respond := &api.ApiResponder{
		Data: paginator,
		Code: 200,
	}

	return respond, err
}

func (ctrl *actionCtrl) Read(id string, r *http.Request) (api.Responder, error) {
	service := ctrl.service
	payload := &entity.Action{}
	payload.SetID(id)
	action, err := service.Read(*payload)

	return &api.ApiResponder{
		Data: action,
		Code: 200,
	}, err
}

func (ctrl *actionCtrl) Replace(id string, r *http.Request) (api.Responder, error) {
	ac := entity.Action{}
	ac.SetID(id)
	payload := entity.Action{}
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

	updated, err := ctrl.service.Replace(ac, payload)

	return &api.ApiResponder{
		Data: updated,
		Code: 200,
	}, err
}

func (ctrl *actionCtrl) Edit(id string, r *http.Request) (api.Responder, error) {
	ac := entity.Action{}
	ac.SetID(id)
	payload := entity.Action{}
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

	updated, err := ctrl.service.Edit(ac, payload)

	return &api.ApiResponder{
		Data: updated,
		Code: 200,
	}, err
}

func (ctrl *actionCtrl) Add(r *http.Request) (api.Responder, error) {
	payload := entity.Action{}
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

	action, err := ctrl.service.Add(payload)

	return &api.ApiResponder{
		Data: action,
		Code: 200,
	}, err
}

func (ctrl *actionCtrl) Delete(id string, r *http.Request) (api.Responder, error) {
	ac := entity.Action{}
	ac.SetID(id)

	_, err := ctrl.service.Delete(ac)

	return &api.ApiResponder{
		Data: nil,
		Code: 204,
	}, err
}

func (ctrl *actionCtrl) BatchAdd(r *http.Request) (api.Responder, error) {
	var payloads []entity.Action
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

func (ctrl *actionCtrl) BatchEdit(r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *actionCtrl) Destroy(r *http.Request) (api.Responder, error) {
	return nil, nil
}
