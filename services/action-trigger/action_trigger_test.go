package actiontrigger

import (
	"fmt"
	"os"
	"testing"

	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/modules/setting"
	"github.com/kwri/go-workflow/services/entity"
	"github.com/stretchr/testify/assert"
)

var (
	service    *ActionTriggerService
	dbfixtures []*entity.ActionTrigger
)

func setup() {
	setting.ConfigFile = "../../config.ini"
	setting.Initialize()
	db.Initialize()
	service = NewActionTriggerService()
	tx := service.Repo.db.Begin()
	service.Repo.db = tx
	seedTestData()
}

func shutdown() {
	service.Repo.db.Rollback()
	db.Engine.Close()
}

func seedTestData() {
	for i := 1; i <= 10; i++ {
		model := &entity.ActionTrigger{
			Min: fmt.Sprintf("test-seed-data-%d", i),
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
	a := &entity.SearchAdapter{}
	result, err := service.Browse(a)
	assert.Nil(t, err, "Error is nil")
	assert.NotEqual(t, result, "Data should not empty")
}

func TestAddActionTrigger(t *testing.T) {
	model := entity.ActionTrigger{
		Min: "test",
	}

	result, err := service.Add(model)

	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, model.Min, result.Min, "It should be equal")
	assert.NotEqual(t, model.ID, result.ID, "Result ID should not be equal")
}

func TestReadActionTrigger(t *testing.T) {
	fixtures := dbfixtures[0]
	model := entity.ActionTrigger{
		ID: fixtures.ID,
	}
	result, err := service.Read(model)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixtures.Min, result.Min, "it should be equal")
	assert.Equal(t, fixtures.ID, result.ID, "it should be equal")
}

func TestEditActionTrigger(t *testing.T) {
	fixture := dbfixtures[0]
	model := entity.ActionTrigger{
		Min: "edited-test-seed-data",
	}
	result, err := service.Edit(*fixture, model)
	if err.Error() != "record not found" {
		assert.Nil(t, err, "Error is nil")
		assert.NotEqual(t, fixture.Min, result.Min, "It should not be equal")
		assert.Equal(t, model.Min, result.Min, "It should be equal")
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
	fixture := entity.ActionTrigger{}
	_, err := service.Delete(fixture)
	assert.NotNil(t, err, "Error is not nil")
}
