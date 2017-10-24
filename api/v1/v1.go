package v1

import (
	api "github.com/kwri/go-workflow/vndapi"
)

func RegisterApi(r *api.Api) {
	r.Resource("workflows", newWorkflowCtrl())
}
