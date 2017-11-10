package v1

import (
	"io/ioutil"
	"net/http"

	paginator "github.com/kwri/go-workflow/gorm-paginator"
	"github.com/kwri/go-workflow/services/entity"
	"github.com/kwri/go-workflow/services/rule"
	api "github.com/kwri/go-workflow/vndapi"
	"github.com/manyminds/api2go/jsonapi"
)

type ruleCtrl struct {
	service *rule.RuleService
}

func newRuleCtrl() *ruleCtrl {
	return &ruleCtrl{
		service: rule.NewRuleService(),
	}
}

func (ctrl *ruleCtrl) Browse(r *http.Request) (api.Responder, error) {
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

	var rules []*entity.Rule

	if total > 0 {
		rules, err = service.Browse(adapter)
	}

	paginator := paginator.NewLengthAwareOffsetPaginator(rules, total, limit, offset, options)

	respond := &api.ApiResponder{
		Data: paginator,
		Code: 200,
	}

	return respond, err
}

func (ctrl *ruleCtrl) Read(id string, r *http.Request) (api.Responder, error) {
	service := ctrl.service
	payload := &entity.Rule{}
	payload.SetID(id)

	rule, err := service.Read(*payload)

	respond := &api.ApiResponder{
		Data: rule,
		Code: 200,
	}

	return respond, err
}

func (ctrl *ruleCtrl) Replace(id string, r *http.Request) (api.Responder, error) {
	rule := entity.Rule{}
	rule.SetID(id)
	payload := entity.Rule{}

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

	updated, err := ctrl.service.Replace(rule, payload)

	respond := &api.ApiResponder{
		Data: updated,
		Code: 200,
	}

	return respond, err
}

func (ctrl *ruleCtrl) Edit(id string, r *http.Request) (api.Responder, error) {
	rule := entity.Rule{}
	rule.SetID(id)
	payload := entity.Rule{}

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

	updated, err := ctrl.service.Edit(rule, payload)

	respond := &api.ApiResponder{
		Data: updated,
		Code: 200,
	}

	return respond, err
}

func (ctrl *ruleCtrl) Add(r *http.Request) (api.Responder, error) {
	payload := entity.Rule{}
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

	rule, err := ctrl.service.Add(payload)

	respond := &api.ApiResponder{
		Data: rule,
		Code: 200,
	}

	return respond, err
}

func (ctrl *ruleCtrl) Delete(id string, r *http.Request) (api.Responder, error) {
	wk := entity.Rule{}
	wk.SetID(id)

	_, err := ctrl.service.Delete(wk)

	return &api.ApiResponder{
		Data: nil,
		Code: 204,
	}, err
}

func (ctrl *ruleCtrl) BatchAdd(r *http.Request) (api.Responder, error) {
	var payloads []entity.Rule
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

func (ctrl *ruleCtrl) BatchEdit(r *http.Request) (api.Responder, error) {
	return nil, nil
}

func (ctrl *ruleCtrl) Destroy(r *http.Request) (api.Responder, error) {
	return nil, nil
}
