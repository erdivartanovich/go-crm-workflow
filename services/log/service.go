package log

type LogService struct {
	Repo *LogRepository
}

type SearchAdapter struct {}


func NewLogService() *LogService {
	return &LogService{
		Repo: NewLogRepository(),
	}
}