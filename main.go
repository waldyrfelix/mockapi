package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/waldyrfelix/mockapi/app"
)

func main() {
	godotenv.Load()

	fmt.Print("Initializing server...")

	app := &app.App{}
	app.SetRouters()
	app.SetMiddlewares()

	fmt.Println("Server on localhost:" + os.Getenv("PORT"))

	app.Run(":" + os.Getenv("PORT"))
}
