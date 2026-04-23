package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	myApp "github.com/lyapkin/search_comg_xls/internal/app"
)

func main() {
	fmt.Println("start app")
	a := app.NewWithID("search_comb_xls")
	fmt.Println("app initialized")
	window := myApp.New(a)
	fmt.Println("app show and run")
	window.ShowAndRun()
}
