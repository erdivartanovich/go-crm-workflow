package action_trigger
import (
	"github.com/stretchr/testify/assert"
	"fmt"
	"os"
	"github.com/kwri/go-workflow/modules/db"
	"github.com/kwri/go-workflow/modules/setting"
	"testing"
 )

var (
	service *ActionTriggerService
	dbfixtures []*ActionTrigger
)

func setup() {
	setting.ConfigFile = "../../config.ini"
	setting.Initialize()
	db.Initialize()
	service = NewActionTriggerService()
	tx := service.Repo.db.Begin()
	service.Repo.db = tx;
	seedTestData()
}

func shutdown() {
	service.Repo.db.Rollback()
	db.Engine.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}