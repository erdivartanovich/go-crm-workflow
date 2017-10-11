package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kwri/go-workflow/api/v1/workflow"
)

func RegisterRoute(r *gin.RouterGroup) {
	wfr := r.Group("/workflows")
	{
		wfr.GET("/", workflow.Browse)
		wfr.GET("/:id", workflow.Read)
		wfr.PATCH("/:id", workflow.Edit)
		wfr.POST("/", workflow.Add)
		wfr.DELETE("/:id", workflow.Delete)
	}
}
