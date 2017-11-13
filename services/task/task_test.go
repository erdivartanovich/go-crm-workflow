package task

import (
	"fmt"
	"os"
	"time"

	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/modules/setting"
	"github.com/kwri/go-workflow/services/entity"
	"github.com/stretchr/testify/assert"

	"testing"
)

var (
	service    *TaskService
	dbfixtures []*entity.Task
)

func setup() {
	setting.ConfigFile = "../../config.ini"
	setting.Initialize()
	db.Initialize()
	service = NewTaskService()
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
		model := &entity.Task{
			UserID:            1,
			TaskType:          4,
			TaskAction:        fmt.Sprintf("test-seed-data-%d", i),
			DueDate:           time.Now(),
			FromInteraction:   fmt.Sprintf("test-seed-data-%d", i),
			Reason:            fmt.Sprintf("test-seed-data-%d", i),
			Description:       fmt.Sprintf("test-seed-data-%d", i),
			IsCompleted:       1,
			IsAutomated:       0,
			CreatedBy:         1,
			UpdatedBy:         1,
			Status:            1,
			MinimumCompletion: 4,
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

func TestAddTask(t *testing.T) {

	model := entity.Task{
		UserID:            1,
		TaskType:          4,
		TaskAction:        "Test",
		DueDate:           time.Now(),
		FromInteraction:   "Test",
		Reason:            "Test",
		Description:       "Test",
		IsCompleted:       1,
		IsAutomated:       0,
		CreatedBy:         1,
		UpdatedBy:         1,
		Status:            1,
		MinimumCompletion: 4,
	}

	result, err := service.Add(model)

	assert.Nil(t, err, "Error is nil")
	assert.NotEqual(t, model.ID, result.ID, "Result ID should have new value")
}

func TestReadTask(t *testing.T) {
	fixture := dbfixtures[0]
	model := entity.Task{
		ID: fixture.ID,
	}
	result, err := service.Read(model)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
}

func TestEditTask(t *testing.T) {
	fixture := entity.Task{
		ID: dbfixtures[0].ID,
	}

	model := entity.Task{
		TaskAction: "edited-test-seed-data",
	}

	result, err := service.Edit(fixture, model)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
}

func TestDeleteTask(t *testing.T) {
	fixture := dbfixtures[0]
	result, err := service.Delete(*fixture)
	assert.Nil(t, err, "Error is nil")
	assert.Equal(t, fixture.ID, result.ID, "It should be equal")
	assert.NotEmpty(t, result.DeletedAt, "Result ID should have new value")
}

func TestDeleteEmptyIDError(t *testing.T) {
	fixture := entity.Task{}
	_, err := service.Delete(fixture)
	assert.NotNil(t, err, "Error is not nil")
}
