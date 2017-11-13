package task

import (
	"github.com/kwri/go-workflow/services/entity"
)

type TaskService struct {
	Repo *TaskRepository
}

func (service *TaskService) Browse(adapter *entity.SearchAdapter) ([]*entity.Task, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *TaskService) Read(task entity.Task) (*entity.Task, error) {
	return service.Repo.Where(task).First()
}

func (service *TaskService) Edit(task entity.Task, payload entity.Task) (*entity.Task, error) {
	wk, err := service.Read(task)

	if err != nil {
		return nil, err
	}
	return service.Repo.Update(*wk, payload)
}

func (service *TaskService) Replace(task entity.Task, payload entity.Task) (*entity.Task, error) {
	return service.Repo.Replace(task, payload)
}

func (service *TaskService) Add(task entity.Task) (*entity.Task, error) {
	return service.Repo.Insert(task)
}

func (service *TaskService) BatchAdd(payloads []entity.Task) (int, error) {
	var ch chan bool
	ch = make(chan bool)

	go func() {
		for i := range payloads {
			_, err := service.Add(payloads[i])
			if err != nil {
				ch <- false
				continue
			}
			ch <- true
		}
		close(ch)
	}()
	success := 0
	for n := range ch {
		if n == true {
			success++
		}
	}

	return success, nil
}

func (service *TaskService) Delete(task entity.Task) (*entity.Task, error) {
	return service.Repo.Delete(task)
}

func (service *TaskService) Count(adapter *entity.SearchAdapter) (int, error) {
	return service.Repo.SetAdapter(adapter).Count()
}

func NewTaskService() *TaskService {
	return &TaskService{
		Repo: NewTaskRepository(),
	}
}
