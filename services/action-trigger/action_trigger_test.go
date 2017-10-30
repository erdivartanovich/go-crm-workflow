package action_trigger
import (
	"github.com/stretchr/testify/assert"
	"fmt"
	"os"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/modules/setting"
	"testing"
 )

var (
	service *ActionTriggerService
	dbfixtures []*ActionTrigger
)

func setup() {
	setting.ConfigFile = "../../config.ini"
	setting.Initialize()
	db.Initialize()
	service = NewActionTriggerService()
	tx := service.Repo.db.Begin()
	service.Repo.db = tx;
	seedTestData()
}

func shutdown() {
	service.Repo.db.Rollback()
	db.Engine.Close()
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
	assert.NotEqual(t, result, "Data should not empty")
}

func TestAddActionTrigger(t, *testing.T) {
	model := ActionTrigger{
		UserID: 1,
		Name: "test",
	}

	result, err := service.Add(model)

	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, mode.Name, result.Name, "It should be equal")
	assert.NotEqual(t, model.ID, result.Id, "Result ID should not be equal")
}