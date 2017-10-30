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

func TestReadActionTrigger(t *testing.T) {
	fixtures := dbfixtures[0]
	model := ActionTrigger{
		ID: fixtures.ID,
	}
	result, err := service.Read(model)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixtures.Name, result.Name, "it should be equal")
	assert.Equal(t, fixtures.ID, result.ID, "it should be equal")
}

func TestEditActionTrigger(t *testing.T) {
	fisture := dbfixtures[0]
	model := ActionTrigger{
		Name: "edited-test-seed-data"
	}
	result, err := service.Edit(*fixture. model)
	if err.Error() != "record not found" {
		assert.Nil(t, err, "Error is nil")
		assert.NotEqual(t, fixture.Name, result.Name, "It should not be equal")
		assert.Equal(t, model.Name, result.Name, "It should be equal")
		assert.Equal(t, fixture.ID, result.ID, "It should be equal")
	} else {
		assert.Equal(t, err.Error(), "record not found", "It should be equal")
	}
}

func TestDeleteActionTrigger(t *testing.T) {
	fixture := dbfixtures[0]
	result, err := service.Delete(*fixture)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
	assert.NotEmpty(t, result.DeletedAt, "Result ID should have new value")
}

func TestDeleteEmptyIDError(t *testing.T) {
	fixture := ActionTrigger{}
	_, err := service.Delete(fixture)
	assert.NotNil(t, err, "Error is not nil")
}