package routes

import (
	"go-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/", controllers.Hello)

	app.Post("/api/register", controllers.Register)

	app.Post("/api/login", controllers.Login)

	app.Get("/api/user", controllers.User)

	app.Post("/api/logout", controllers.Logout)
}
