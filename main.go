package main

import (
	"fmt"
	"os"

	"github.com/waldyrfelix/mockapi/app"
)

func main() {
	fmt.Print("Initializing server...")
	app := &app.App{}
	app.SetRouters()
	app.SetMiddlewares()

	fmt.Println("Server on localhost:3000.")

	app.Run(":" + os.Getenv("PORT"))
}
