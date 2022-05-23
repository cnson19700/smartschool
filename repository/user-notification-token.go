package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

type UserNotificationTokenRepository interface {
	GetUserIDArray(uint) []uint
}

func GetUserIDArray(notificationID uint) (userIDArray []uint) {
	userNotiTokenList := make([]entity.UserNotificationToken, 0)
	database.DbInstance.
		Where("notification_id = ?", notificationID).
		Find(&userNotiTokenList)

	for _, value := range userNotiTokenList {
		userIDArray = append(userIDArray, value.UserID)
	}

	return userIDArray
}
