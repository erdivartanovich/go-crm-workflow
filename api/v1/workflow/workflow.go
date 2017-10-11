package workflow

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwri/go-workflow/services/workflow"
)

func Browse(c *gin.Context) {
	service := workflow.NewWorkflowService()
	adapter := workflow.SearchAdapter{}
	result, err := service.Browse(adapter)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, result)
}

func Read(c *gin.Context) {

}

func Edit(c *gin.Context) {

}

func Add(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
