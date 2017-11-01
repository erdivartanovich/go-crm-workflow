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

func TestAddLog(t *testing.T) {

	model := entity.WorkflowLog{
		ResourceName: "test",
		UserID: 1,
		Info: "test info",
	}

	result, err := service.Add(model)

	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, model.ResourceName, result.ResourceName, "It should be equal")
	assert.NotEqual(t, model.ID, result.ID, "Result ID should have new value")
}

func TestReadLog(t *testing.T) {
	fixture := dbfixtures[0]
	model := entity.WorkflowLog{
		ID: fixture.ID,
	}
	result, err := service.Read(model)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixture.Name, result.Name, "It should be equal")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
}

func TestEditLog(t *testing.T) {
	fixture := entity.WorkflowLog{
		ID: dbfixtures[0].ID,
	}

	model := entity.WorkflowLog{
		Name: "edited-test-seed-data",
	}

	result, err := service.Edit(fixture, model)
	assert.Nil(t, err, "Error is nil")
	assert.NotEqual(t, fixture.ResourceName, result.ResourceName, "It should not be equal")
	assert.Equal(t, model.ResourceName, result.ResourceName, "It should be equal")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
}

func TestDeleteLog(t *testing.T) {
	fixture := dbfixtures[0]
	result, err := service.Delete(*fixture)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
	assert.NotEmpty(t, result.DeletedAt, "Result ID should have new value")
}

func TestDeleteEmptyIDError(t *testing.T) {
	fixture := entity.WorkflowLog{}
	_, err := service.Delete(fixture)
	assert.NotNil(t, err, "Error is not nil")
}