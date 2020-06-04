package main

import (
	"MYGO/studying/dependency-injection/controller"
	"MYGO/studying/dependency-injection/store"
	"MYGO/studying/dependency-injection/svc"
	"fmt"
	"log"

	"go.uber.org/dig"
)

func main() {
	fmt.Println("Start inject ====>")
	c := dig.New()

	// err := c.Provide(store.NewMessage)
	// if err != nil {
	// 	log.Fatalln("Can not inject store", err.Error())
	// }

	err := c.Provide(func() store.Message {
		return store.Message("This is new messaging!!")
	})
	if err != nil {
		log.Fatalln("Can not inject store", err.Error())
	}

	err = c.Provide(svc.NewGreeter)
	if err != nil {
		log.Fatalln("Can not inject svc", err.Error())
	}

	err = c.Provide(controller.NewEvent)
	if err != nil {
		log.Fatalln("Can not inject controller", err.Error())
	}

	c.Invoke(controller.EventCtrl.PushMessage)

	fmt.Println("End inject ====>")
}
