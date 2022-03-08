package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

var Generators = map[string]table.Generator{
	"courses": GetCourses,
	"summary": GetSummary,
}