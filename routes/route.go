package routes

import (
	"apiant/controllers"
	"apiant/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Routes(r *fiber.App) {
	ant := r.Group("/api")

	ant.Get("/", middlewares.AuthMiddleware, controllers.Index)
	ant.Get("/:id", middlewares.AuthMiddleware, controllers.Show)
	ant.Post("/:id", middlewares.AuthMiddleware, controllers.Update)
	ant.Post("/reset/:id", middlewares.AuthMiddleware, controllers.Reset)
}