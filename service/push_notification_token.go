package service

import (
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)

func PushNotificationForAdminAndTeacher(registerID uint, courseID uint, user entity.User, notificationTitle string, notificationAction string, notificationMessage string) error {
	var adminMap = make(map[uint]*entity.User)
	//receiverArray := make([]*entity.User, 0)
	teachersMap := make([]entity.Teacher, 0)

	// Find all admin of department
	var adminOfDepartmentArray = make([]*entity.User, 0)
	adminOfDepartmentArray, err := repository.QueryAdminByFacultyID(user.FacultyID)
	if err != nil {
		return err
	}

	for _, admin := range adminOfDepartmentArray {
		adminMap[admin.ID] = admin
	}

	var teacherArray = make([]*entity.Teacher, 0)
	teacherArray, err = repository.QueryTeacherIDByCourseID(courseID)

	for _, teacher := range teacherArray {
		teachersMap[teacher.ID] = *teacher
	}

	// for _, admin := range adminMap {
	// 	receiverArray = append(receiverArray, teacherArray...)
	// }

	return nil
	//worker.SendNewRegisterNotification(registerID, receiverArray, notificationTitle, notificationAction, notificationMessage)
}
