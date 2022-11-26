package main

import (
	"fmt"
	"os"

	"github.com/uekiGityuto/othello-go/app/controller"
	"github.com/uekiGityuto/othello-go/app/model"
)

func main() {
	ctrl, err := controller.New(model.White)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ctrl.Start()
}
