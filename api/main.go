package main

import (
	"fmt"

	// "github.com/smartschool/api/routers"
	"os"
	"os/signal"

	"github.com/smartschool/api/routers"
	"github.com/smartschool/database"
)

func main() {
	fmt.Println("=================================")
	fmt.Println("Start bSmartCheckin Core API......")
	fmt.Println("=================")

	database.Init()

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
