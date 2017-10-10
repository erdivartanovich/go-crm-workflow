package generator

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/kwri/go-workflow/modules/fs"
	"github.com/kwri/go-workflow/modules/logger"
)

type Service struct {
	Name string
}

func NewService(name string) Service {
	return Service{
		Name: name,
	}
}

func (service Service) Generate() {
	service.createDirectory()
	service.createService()
	service.createEntity()
	service.createRepository()
}

func (service Service) createDirectory() {
	dirname := "./services/" + service.Name
	if !fs.FileExists(dirname) {
		err := fs.CreateDirectory(dirname)
		if err != nil {
			logger.Panic(err)
		}
	}
}

func (service Service) createService() {
	err := fs.FilePutContent(service.serviceName(), service.serviceTempl())
	if err != nil {
		logger.Panic(err)
	}
}

func (service Service) createEntity() {
	err := fs.FilePutContent(service.entityName(), service.entityTempl())
	if err != nil {
		logger.Panic(err)
	}
}

func (service Service) createRepository() {
	err := fs.FilePutContent(
		service.repositoryName(),
		service.repositoryTempl(),
	)
	if err != nil {
		logger.Panic(err)
	}
}

func (service Service) serviceName() string {
	return "./services/" + service.Name + "/service.go"
}

func (service Service) entityName() string {
	return "./services/" + service.Name + "/entity.go"
}

func (service Service) repositoryName() string {
	return "./services/" + service.Name + "/repository.go"
}

func (service Service) serviceTempl() string {
	return fmt.Sprintf(serviceTempl, service.Name)
}

func (service Service) entityTempl() string {
	return fmt.Sprintf(
		entityTempl,
		service.Name,
		strcase.ToCamel(service.Name),
		addOrmQuote(`gorm:"primary_key"`),
		addOrmQuote(`gorm:"unsigned user_id;unique_index:name_user_id"`),
		addOrmQuote(`gorm:"not null;unique_index:name_user_id"`),
		addOrmQuote(`gorm:"default:current_timestamp"`),
		addOrmQuote(`gorm:"default:current_timestamp on update current_timestamp"`),
	)
}

func (service Service) repositoryTempl() string {
	return fmt.Sprintf(repositoryTempl, service.Name)
}

func addOrmQuote(str string) string {
	return "`" + str + "`"
}
