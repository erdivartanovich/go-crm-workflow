package log

import (
	"fmt"
	"os"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/modules/setting"
	"github.com/kwri/go-workflow/services/entity"
)

var (
	service *LogService
	dbfixtures []*entity.WorkflowLog
)

func setup() {
	setting.ConfigFile = "../../config.ini"
	setting.Initialize()
	db.Initialize()
	service = NewLogService()
	tx := service.Repo.db.Begin()
	service.Repo.db = tx
	seedTestData()
}