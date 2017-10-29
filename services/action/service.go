package action

type ActionService struct {
	Repo *ActionRepository
}

type SearchAdapter struct {
}

func (service *ActionService) Browse(adapter SearchAdapter) ([]*Action, error) {
	return service.Repo.SetAdapter(adapter).Find()
}

func (service *ActionService) Read(action Action) (*Action, error) {
	return service.Repo.Where(action).First()
}

func (service *ActionService) Edit(action Action, payload Action) (*Action, error) {
	return service.Repo.Update(action, payload)
}

func (service *ActionService) Add(action Action) (*Action, error) {
	return service.Repo.Insert(action)
}

func (service *ActionService) Delete(action Action) (*Action, error) {
	return service.Repo.Delete(action)
}

func (service *ActionService) Count(adapter SearchAdapter) (int, error) {
	return service.Repo.SetAdapter(adapter).Count()
}

func NewActionService() *ActionService {
	return &ActionService{
		Repo:NewActionRepository(),
	}
}