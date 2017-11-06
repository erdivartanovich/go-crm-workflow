package log

import (
	"github.com/kwri/go-workflow/services/entity"
)
type LogService struct {
	Repo *LogRepository
}

type SearchAdapter struct {
}


func (service *LogService) Browse(adapter SearchAdapter) ([]*entity.WorkflowLog, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *LogService) Read(workflowlog entity.WorkflowLog) (*entity.WorkflowLog, error) {
	return service.Repo.Where(workflowlog).First()
}

func (service *LogService) Edit(workflowlog entity.WorkflowLog, payload entity.WorkflowLog) (*entity.WorkflowLog, error) {
	wl, err := service.Read(workflowlog)

	if err != nil {
		return nil, err
	}
	return service.Repo.Update(*wl, payload)
}

func (service *LogService) Replace(workflowlog entity.WorkflowLog, payload entity.WorkflowLog) (*entity.WorkflowLog, error) {
	return service.Repo.Replace(workflowlog, payload)
}

func (service *LogService) Add(workflowlog entity.WorkflowLog) (*entity.WorkflowLog, error) {
	return service.Repo.Insert(workflowlog)
}

func (service *LogService) BatchAdd(payloads []entity.WorkflowLog) (int, error) {
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

func (service *LogService) Delete(workflowlog entity.WorkflowLog) (*entity.WorkflowLog, error) {
	return service.Repo.Delete(workflowlog)
}

func (service *LogService) Count(adapter SearchAdapter) (int, error) {
	return service.Repo.SetAdapter(adapter).Count()
}

func NewLogService() *LogService {
	return &LogService{
		Repo: NewLogRepository(),
	}
}