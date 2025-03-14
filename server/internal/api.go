package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func NewApi() *fiber.App {
	// создаём объект, который оперирует с пользователями в базе данных
	userSt = CreateUserStorageDb()
	deviceSt = CreateDeviceStorageDb()

	app := fiber.New()

	app.Use(cors.New( cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Access-Control-Allow-Origin, Authorization",
		AllowOrigins:     "http://localhost:3000, http://84.201.140.80",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Static("/static", "./static")

	// подключаем маршруты для нашего приложения
	AttachRoutes(app)

	return app
}