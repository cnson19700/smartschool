package fireapp

import (
	"errors"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App

func Init() error {
	opt := option.WithCredentialsFile("resources/smartschool-dev-aeed1-firebase-adminsdk-q0mkq-13527cab01.json")

	var err error
	firebaseApp, err = firebase.NewApp(context.Background(), nil, opt)
	return err
}

func SendNotification(userId uint, data map[string]string) error {
	var tokenObjs []entity.UserNotificationToken
	err := database.DbInstance.Where("user_id = ?", userId).Find(&tokenObjs).Error
	if err != nil {
		return err
	}
	if len(tokenObjs) == 0 {
		return errors.New("no notification tokens found for user")
	}

	tokens := make([]string, 0)
	for _, tokenObj := range tokenObjs {
		tokens = append(tokens, tokenObj.Token)
	}

	message := &messaging.MulticastMessage{
		Data:   data,
		Tokens: tokens,
	}

	client, err := firebaseApp.Messaging(context.Background())
	if err != nil {
		return err
	}

	responses, err := client.SendMulticast(context.Background(), message)
	if err != nil {
		return err
	}
	if responses.FailureCount > 0 {
		return errors.New("failure sending notification")
	}

	return nil
}
