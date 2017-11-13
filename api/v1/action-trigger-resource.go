package v1

import (
	"io/ioutil"
	"net/http"

	paginator "github.com/kwri/go-workflow/gorm-paginator"
	actiontrigger "github.com/kwri/go-workflow/services/action-trigger"
	"github.com/kwri/go-workflow/services/entity"
	api "github.com/kwri/go-workflow/vndapi"
	"github.com/manyminds/api2go/jsonapi"
)

type actionTriggerCtrl struct {
	service *actiontrigger.ActionTriggerService
}

func newActionTriggerCtrl() *actionTriggerCtrl {
	return &actionTriggerCtrl{
		service: actiontrigger.NewActionTriggerService(),
	}
}

func (ctrl *actionTriggerCtrl) Browse(r *http.Request) (api.Responder, error) {
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

	var actionTriggers []*entity.ActionTrigger

	if total > 0 {
		actionTriggers, err = service.Browse(adapter)
	}

	paginator := paginator.NewLengthAwareOffsetPaginator(actionTriggers, total, limit, offset, options)

	respond := &api.ApiResponder{
		Data: paginator,
		Code: 200,
	}

	return respond, err
}

func (ctrl *actionTriggerCtrl) Read(id string, r *http.Request) (api.Responder, error) {
	service := ctrl.service
	payload := &entity.ActionTrigger{}
	payload.SetID(id)
	actionTrigger, err := service.Read(*payload)

	return &api.ApiResponder{
		Data: actionTrigger,
		Code: 200,
	}, err
}

func (ctrl *actionTriggerCtrl) Replace(id string, r *http.Request) (api.Responder, error) {

	wk := entity.ActionTrigger{}
	wk.SetID(id)
	payload := entity.ActionTrigger{}
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

func (ctrl *actionTriggerCtrl) Edit(id string, r *http.Request) (api.Responder, error) {
	wk := entity.ActionTrigger{}
	wk.SetID(id)
	payload := entity.ActionTrigger{}
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

func (ctrl *actionTriggerCtrl) Add(r *http.Request) (api.Responder, error) {
	payload := entity.ActionTrigger{}
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

	actionTrigger, err := ctrl.service.Add(payload)
	//
	return &api.ApiResponder{
		Data: actionTrigger,
		Code: 200,
	}, err
}

func (ctrl *actionTriggerCtrl) Delete(id string, r *http.Request) (api.Responder, error) {
	wk := entity.ActionTrigger{}
	wk.SetID(id)

	_, err := ctrl.service.Delete(wk)

	return &api.ApiResponder{
		Data: nil,
		Code: 204,
	}, err
}

func (ctrl *actionTriggerCtrl) BatchAdd(r *http.Request) (api.Responder, error) {
	var payloads []entity.ActionTrigger
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

func (ctrl *actionTriggerCtrl) BatchEdit(r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *actionTriggerCtrl) Destroy(r *http.Request) (api.Responder, error) {
	return nil, nil
}
