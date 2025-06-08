package routes

import (
	"CambodiaTaxPortal/app/http/controllers"

	"github.com/goravel/framework/facades"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)
}
