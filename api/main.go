package main

import (
	"fmt"
	"github.com/smartschool/api/routers"
	"github.com/smartschool/database"
	"github.com/smartschool/service/fireapp"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("=================================")
	fmt.Println("Start bSmartCheckin Core API......")
	fmt.Println("=================")

	database.Init()

	defer database.Close()

	err := fireapp.Init()
	if err != nil {
		fmt.Println("Firebase error: " + err.Error())
		return
	}

	go func() {
		r, err := routers.Initialize()
		if err != nil {
			fmt.Println(err)
		} else {
			r.Run(":6002")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	fmt.Println("=======================")
	fmt.Println("bSmartCheckin Core API is closing... See ya!")
}
