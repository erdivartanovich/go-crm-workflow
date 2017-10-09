package migrations

import (
	"github.com/kwri/go-workflow/modules/migrate"
)

var (
	Migrations = []*migrate.Migration{
		&create_table_workflows_1507565382,
	}
)
