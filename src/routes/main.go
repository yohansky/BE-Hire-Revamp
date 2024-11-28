package routes

import (
	"be-hire-revamp/src/controllers"
	"be-hire-revamp/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/register-worker", controllers.RegisterWorker)
	app.Post("/register-recruiter", controllers.RegisterRecruiter)
	app.Post("/login", controllers.Login)

	app.Use(middleware.IsAuth)

	app.Put("/user/info", controllers.UpdateInfor)
	app.Put("/user/password", controllers.UpdatePassword)

	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)

	app.Get("/roles", controllers.AllRoles)
	app.Post("/roles", controllers.CreateRole)
	app.Get("/role/:id", controllers.GetRole)
	app.Delete("/role/:id", controllers.DeleteRole)
}
