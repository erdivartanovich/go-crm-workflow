package action

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
	service    *ActionService
	dbfixtures []*entity.Action
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
	for i := 1; i <= 10; i++ {
		model := &entity.Action{
			Name:   fmt.Sprintf("test-seed-data-%d", i),
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
	a := entity.SearchAdapter{}
	result, err := service.Browse(a)
	assert.Nil(t, err, "Error is nil")
	assert.NotEmpty(t, result, "Data should not empty")
}

func TestAddAction(t *testing.T) {
	model := entity.Action{
		UserID: 1,
		Name:   "test",
	}

	result, err := service.Add(model)

	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, model.Name, result.Name, "It should be equal")
	assert.NotEqual(t, model.ID, result.ID, "Result ID should have new value")
}

func TestReadAction(t *testing.T) {
	fixture := dbfixtures[0]
	model := entity.Action{
		ID: fixture.ID,
	}
	result, err := service.Read(model)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixture.Name, result.Name, "it should be equal")
	assert.Equal(t, fixture.ID, result.ID, "it should be equal")
}

func TestEditAction(t *testing.T) {
	fixture := dbfixtures[0]
	model := entity.Action{
		Name: "edited-test-seed-data",
	}
	result, err := service.Edit(*fixture, model)
	if err.Error() != "record not found" {
		assert.Nil(t, err, "Error is nil")
		assert.NotEqual(t, fixture.Name, result.Name, "It should not be equal")
		assert.Equal(t, model.Name, result.Name, "It should be equal")
		assert.Equal(t, fixture.ID, result.ID, "It should be equal")
	} else {
		assert.Equal(t, err.Error(), "record not found", "It should show not found if empty record")
	}
}

func TestDeleteAction(t *testing.T) {
	fixture := dbfixtures[0]
	result, err := service.Delete(*fixture)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
	assert.NotEmpty(t, result.DeletedAt, "Result ID should have new value")
}

func TestDeleteEmptyIDError(t *testing.T) {
	fixture := entity.Action{}
	_, err := service.Delete(fixture)
	assert.NotNil(t, err, "Error is not nil")
}
