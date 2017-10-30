package migrations

import (
	"github.com/kwri/go-workflow/modules/migrate"
)

var (
	Migrations = []*migrate.Migration{
		&create_table_workflows_1507565382,
		&create_table_actions_1507579635,
		&create_table_rules_1507635905,
		&create_workflow_log_table_1509349752,
	}
)
