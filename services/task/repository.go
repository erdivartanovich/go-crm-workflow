package task

import (
	"errors"

	"github.com/golang-collections/collections/stack"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/services/entity"
)

type TaskRepository struct {
	db      *gorm.DB
	adapter *entity.SearchAdapter
	where   *stack.Stack
}

func (repo *TaskRepository) SetAdapter(adapter *entity.SearchAdapter) *TaskRepository {
	repo.adapter = adapter
	return repo
}

func (repo *TaskRepository) Find() ([]*entity.Task, error) {
	task := &[]*entity.Task{}
	err := repo.prepareDb().Find(task).Error
	repo.ResetInstance()
	return *task, err
}

func (repo *TaskRepository) Where(task entity.Task) *TaskRepository {
	repo.where.Push(&task)
	return repo
}

func (repo *TaskRepository) First() (*entity.Task, error) {
	task := &entity.Task{}
	err := repo.prepareDb().First(task).Error
	repo.ResetInstance()
	return task, err
}

func (repo *TaskRepository) Update(task entity.Task, payload entity.Task) (*entity.Task, error) {
	err := repo.prepareDb().Model(&task).Update(payload).Error
	repo.ResetInstance()
	return &task, err
}

func (repo *TaskRepository) Replace(task entity.Task, payload entity.Task) (*entity.Task, error) {
	wk := &task
	db := repo.prepareDb()
	db.First(wk)
	wk.TaskType = payload.TaskType
	wk.TaskAction = payload.TaskAction
	wk.DueDate = payload.DueDate
	wk.FromInteraction = payload.FromInteraction
	wk.Reason = payload.Reason
	wk.Description = payload.Description
	wk.IsCompleted = payload.IsCompleted
	wk.IsAutomated = payload.IsAutomated
	wk.CreatedBy = payload.CreatedBy
	wk.UpdatedBy = payload.UpdatedBy
	wk.Status = payload.Status
	wk.MinimumCompletion = payload.MinimumCompletion
	err := db.Save(wk).Error
	repo.ResetInstance()
	return wk, err
}

func (repo *TaskRepository) Insert(task entity.Task) (*entity.Task, error) {

	in := &task
	err := repo.prepareDb().Create(in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *TaskRepository) Delete(task entity.Task) (*entity.Task, error) {

	in := &task
	if len(in.ID) == 0 {
		return nil, errors.New("You need to set ID of deleted task")
	}
	err := repo.prepareDb().Delete(&in).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Unscoped().Find(&in).Error
	repo.ResetInstance()
	return in, err
}

func (repo *TaskRepository) prepareDb() *gorm.DB {
	count := repo.where.Len()
	tx := repo.db
	for i := 0; i < count; i++ {
		tx = tx.Where(repo.where.Pop())
	}
	if repo.adapter != nil {
		tx = repo.adapter.ApplySearchAdapter(tx)
	}
	return tx
}

func (repo *TaskRepository) ResetInstance() {
	repo.adapter = nil
}

func (repo *TaskRepository) Count() (int, error) {
	count := 0
	tx := repo.prepareDb().Limit(-1).Offset(-1)
	err := tx.Model(&entity.Task{}).Count(&count).Error
	return count, err
}

func NewTaskRepository() *TaskRepository {
	db := db.Engine
	return &TaskRepository{
		db:    db,
		where: &stack.Stack{},
	}
}
