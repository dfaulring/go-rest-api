package main

import (
	"fmt"
	"net/http"

	"github.com/dfaulring/go-rest-api/internal/database"
	transportHTTP "github.com/dfaulring/go-rest-api/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up our app")

	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("GO REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
