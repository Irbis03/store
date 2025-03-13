package internal

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DeviceStorage interface {
	Create(name string, price int, img string) (Device, error)
	GetAll() ([]Device, error)
	GetOne(id int) (Device, error)
	
}

var deviceSt DeviceStorage

func attachDeviceRouter(apiRouter fiber.Router) {
	router := apiRouter.Group("/device")
	
	router.Post("/", roleMiddleware("ADMIN"),  createDeviceHandler)
	router.Get("/", getAllDeviceHandler)
	router.Get("/:id", getOneDeviceHandler)
}

func createDeviceHandler(c *fiber.Ctx) error {
	// разбираем запрос
	type Request struct {
		Name string `json:"name"`
		Price int `json:"price"`
		Img string `json:"img"`
	}
	
	objRequest := new(Request)

	objRequest.Name = c.FormValue("name")

	priceStr := c.FormValue("price")
	price, err := strconv.Atoi(priceStr) 
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "incorrect price"})
	}
	objRequest.Price = price
	
	// сохраняем фото
	img, err := c.FormFile("img")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error with uploading file"})
	}

	imgName := genImgName()

	if err := c.SaveFile(img, "./static/" + imgName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't save device"})
	}

	// создаём запись в хранилище
	device, err := deviceSt.Create(objRequest.Name, objRequest.Price, imgName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't save device"})
	}

	return c.Status(fiber.StatusCreated).JSON(device)
}

func getAllDeviceHandler(c *fiber.Ctx) error {
	devices, err := deviceSt.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't get devices"})
	}

	return c.Status(fiber.StatusOK).JSON(devices)
}

func getOneDeviceHandler(c *fiber.Ctx) error {
	// получаем id пользователя
	id, err := c.ParamsInt("id", 1)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}
	// TODO: проверить существование девайса
	// извлекаем запись об устройстве из хранилища
	device, err := deviceSt.GetOne(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't get device"})
	}

	return c.Status(fiber.StatusOK).JSON(device)
}

func genImgName() string {
	return uuid.New().String() + ".jpg"
}
