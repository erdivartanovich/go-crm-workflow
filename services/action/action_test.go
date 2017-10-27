package action

import (
	"github.com/stretchr/testify/assert"
	"fmt"
	"os"

	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/modules/setting"
	"testing"
)

var (
	service *ActionService
	dbfixtures []*Action
)

func setup() {
	setting.ConfigFile = "../../config.ini"
	setting.Initialize()
	db.Initialize()
	service = NewActionService()
	tx := service.Repo.db.Begin()
	service.Repo.db = tx
	seedTestData()
}

func shutdown() {
	service.Repo.db.Rollback()
	db.Engine.Close()
}

func seedTestData() {
	for i := 1; i<=10; i++ {
		model := &Action{
			Name: fmt.Sprintf("test-seed-data-%d", i),
			UserID: 1,
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
