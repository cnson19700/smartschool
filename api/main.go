package main

import (
	"fmt"
	"github.com/smartschool/api/routers"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("=================================")
	fmt.Println("Start bSmartCheckin Core API......")
	fmt.Println("=================")

	//model.Initialize()

	go func() {
		r := routers.Initialize()
		r.Run(":6969")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	fmt.Println("=======================")
	fmt.Println("bSmartCheckin Core API is closing... See ya!")
}
