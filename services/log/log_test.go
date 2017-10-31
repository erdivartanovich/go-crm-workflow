package log

import (
	"os"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
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

func shutdown() {
	service.Repo.db.Rollback()
	db.Engine.Close()
}

func seedTestData() {
	fmt.Println("testtttt")
	for i := 1; i <= 10; i++ {
		model := &entity.WorkflowLog{
			ResourceName: fmt.Sprintf("test-seed-data-%d", i),
			UserID: 1,
			Info: fmt.Sprintf("test-seed-data-info-%d", i),
		}
		service.Repo.db.Create(model)
		dbfixtures = append(dbfixtures, model)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestBrowse(t *testing.T) {
	a := SearchAdapter{}
	result, err := service.Browse(a)
	assert.Nil(t, err, "Error is nil")
	assert.NotEmpty(t, result, "Data should not empty")
}