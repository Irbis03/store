package internal

import "github.com/gofiber/fiber/v2"

func AttachRoutes(app *fiber.App) {
	apiRouter := app.Group("/api")

	attachUserRouter(apiRouter)
	attachDeviceRouter(apiRouter)
}