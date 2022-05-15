package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

var Generators = map[string]table.Generator{
	"courses":         GetCourses,
	"summary":         GetSummary,
	"users":           GetUsers,
	"attendances":     GetAttendances,
	"teachers":        GetTeachers,
	"teacher_courses": GetTeacherCourses,
}
