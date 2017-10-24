package api

import (
	"fmt"

	"github.com/kwri/go-workflow/api/v1"
	"github.com/kwri/go-workflow/modules/setting"
	api "github.com/kwri/go-workflow/vndapi"
)

func RegisterApi() {
	app := api.New("v1")
	v1.RegisterApi(app)
	api.ListenAndServe(fmt.Sprintf("%s:%s", setting.Config.ApiHost, setting.Config.ApiPort))
}
