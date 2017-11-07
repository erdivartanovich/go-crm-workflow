package v1


import (
	"io/ioutil"
	"net/http"
	"strconv"

	paginator "github.com/kwri/go-workflow/gorm-paginator"
	"github.com/kwri/go-workflow/services/entity"
	"github.com/kwri/go-workflow/services/action"
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
	adapter := action.SearchAdapter{}
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