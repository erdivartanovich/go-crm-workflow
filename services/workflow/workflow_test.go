package workflow

import (
	"fmt"
	"os"

	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/modules/setting"
	"github.com/kwri/go-workflow/services/entity"
	"github.com/stretchr/testify/assert"

	"testing"
)

var (
	service    *WorkflowService
	dbfixtures []*entity.Workflow
)

func setup() {
	setting.ConfigFile = "../../config.ini"
	setting.Initialize()
	db.Initialize()
	service = NewWorkflowService()
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
		model := &entity.Workflow{
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
	a := &entity.SearchAdapter{}
	result, err := service.Browse(a)
	assert.Nil(t, err, "Error is nil")
	assert.NotEmpty(t, result, "Data should not empty")
}

func TestAddWorkflow(t *testing.T) {

	model := entity.Workflow{
		Name:   "test",
		UserID: 1,
	}

	result, err := service.Add(model)

	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, model.Name, result.Name, "It should be equal")
	assert.NotEqual(t, model.ID, result.ID, "Result ID should have new value")
}

func TestReadWorkflow(t *testing.T) {
	fixture := dbfixtures[0]
	model := entity.Workflow{
		ID: fixture.ID,
	}
	result, err := service.Read(model)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixture.Name, result.Name, "It should be equal")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
}

func TestEditWorkflow(t *testing.T) {
	fixture := entity.Workflow{
		ID: dbfixtures[0].ID,
	}

	model := entity.Workflow{
		Name: "edited-test-seed-data",
	}

	result, err := service.Edit(fixture, model)
	assert.Nil(t, err, "Error is nil")
	assert.NotEqual(t, fixture.Name, result.Name, "It should not be equal")
	assert.Equal(t, model.Name, result.Name, "It should be equal")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
}

func TestDeleteWorkflow(t *testing.T) {
	fixture := dbfixtures[0]
	result, err := service.Delete(*fixture)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
	assert.NotEmpty(t, result.DeletedAt, "Result ID should have new value")
}

func TestDeleteEmptyIDError(t *testing.T) {
	fixture := entity.Workflow{}
	_, err := service.Delete(fixture)
	assert.NotNil(t, err, "Error is not nil")
}
