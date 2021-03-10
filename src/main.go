package main

import (
	"fmt"

	"github.com/waldyrfelix/mockapi/src/app"
)

func main() {
	fmt.Print("Initializing server...")
	app := &app.App{}
	app.SetRouters()
	app.SetMiddlewares()
	fmt.Println("Server on localhost:3000.")
	app.Run(":3000")
}
