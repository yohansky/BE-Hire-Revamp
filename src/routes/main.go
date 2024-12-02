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

	app.Get("/users", controllers.AllUsers)
	app.Post("/users", controllers.CreateUser)
	app.Get("/user/:id", controllers.GetUser)
	app.Put("/user/:id", controllers.UpdateUser)
	app.Delete("/user/:id", controllers.DeleteUser)

	app.Get("/roles", controllers.AllRoles)
	app.Post("/roles", controllers.CreateRole)
	app.Get("/role/:id", controllers.GetRole)
	app.Delete("/role/:id", controllers.DeleteRole)

	app.Get("/workers", controllers.AllWorkers)
	app.Post("/workers", controllers.CreateWorker)
	app.Get("/worker/:id", controllers.GetWorker)
	//buat kondisi ketika sudah ada userid akan menemukan workerid
	app.Get("/user/:id/worker", controllers.GetWorkerByUserID)
	app.Put("/worker/:id", controllers.UpdateWorker)
	app.Delete("/worker/:id", controllers.DeleteWorker)

	app.Get("/recruiters", controllers.AllRecruiters)
	app.Post("/recruiters", controllers.CreateRecruiter)
	app.Get("/recruiter/:id", controllers.GetRecruiter)
	app.Put("/recruiter/:id", controllers.UpdateRecruiter)
	app.Delete("/recruiter/:id", controllers.DeleteRecruiter)

	app.Get("/skills", controllers.AllSkills)
	app.Post("/skills", controllers.CreateSkill)
	app.Get("/skill/:id", controllers.GetSkill)
	app.Put("/skill/:id", controllers.UpdateSkill)
	app.Delete("/skill/:id", controllers.DeleteSkill)

	app.Get("/projects", controllers.AllProjects)
	app.Post("/projects", controllers.CreateProject)
	app.Get("/project/:id", controllers.GetProject)
	app.Get("/worker/:id/project", controllers.GetWorkerByWorkerIDProject)
	app.Get("/worker/:id/projects", controllers.GetProjectsByWorkerID)
	app.Put("/project/:id", controllers.UpdateProject)
	app.Delete("/project/:id", controllers.DeleteProject)

	app.Get("/experiences", controllers.AllExperiences)
	app.Post("/experiences", controllers.CreateExperience)
	app.Get("/experience/:id", controllers.GetExperience)
	app.Get("/worker/:id/experience", controllers.GetWorkerByWorkerIDExperience)
	app.Get("/worker/:id/experiences", controllers.GetExperiencesByWorkerID)
	app.Put("/experience/:id", controllers.UpdateExperience)
	app.Delete("/experience/:id", controllers.DeleteExperience)
}
