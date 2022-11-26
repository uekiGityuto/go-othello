package main

import (
	"fmt"

	"github.com/uekiGityuto/othello-go/app/controller"
	"github.com/uekiGityuto/othello-go/app/model"
)

func main() {
	ctrl, err := controller.New(model.White)
	if err != nil {
		fmt.Println(err.Error())
	}
	ctrl.Start()
}
