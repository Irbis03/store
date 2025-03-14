package internal

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	PutUser(email, password, role string) (User, error)
	GetUserByEmail(email string) (User, error)

	// // aditional
	// UpdateUser(user User) error
	// DeleteUser(user_id int) error
	GetAllUser() ([]User, error)
}

var userSt UserStorage
var jwtSecret = []byte("my_secrete_key")

// router

func attachUserRouter(apiRouter fiber.Router) {
	userRouter := apiRouter.Group("/user")
	
	userRouter.Post("/registration", registrationHandler)
	userRouter.Post("/login", loginHandler)
	userRouter.Get("/auth", authMiddleware, authHandler)
	
	// additional
	userRouter.Get("/getAll", getAllHandler)
	// userRouter.Delete("/delete/:id", jwtMiddleware, deleteHandler)
	// userRouter.Put("/update/:id", jwtMiddleware, updateHandler)
}

func registrationHandler(c *fiber.Ctx) error {
	// разбираем запрос
	type Request struct {
		Email string `json:"email"`
		Password string `json:"password"`
		Role string `json:"role"`
	}
	
	objRequest := new(Request)
	if err := c.BodyParser(objRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input" + err.Error()})
	}

	// существует ли пользователь
	if _, exists := userSt.GetUserByEmail(objRequest.Email); exists == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user with provided email exists"})
	}

	// хэшируем пароль
	encryptedPasswrod, err := bcrypt.GenerateFromPassword([]byte(objRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't create user"})
	}

	// сохраняем нового пользователя в базе данных
	user, err := userSt.PutUser(objRequest.Email, string(encryptedPasswrod), objRequest.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't create user"})
	}

	// генерируем JWT token
	tokenStr, err := genJwtToken(user.Id, user.Email, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": tokenStr})
}

func loginHandler(c* fiber.Ctx) error {
	// разбираем запрос
	type Request struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	
	objRequest := new(Request)
	if err := c.BodyParser(objRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input " + err.Error()})
	}

	// существует ли пользователь
	user, exists := userSt.GetUserByEmail(objRequest.Email)
	if exists != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user with provided email doesn't exist"})
	}
	
	// сравниваем пароли
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(objRequest.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// генерируем jwt token
	tokenStr, err := genJwtToken(user.Id, user.Email, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't login"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenStr})
}

func authMiddleware(c *fiber.Ctx) error {
	// Проверяем заголовок c токеном
	authHeader := c.Get("Authorization")
	if authHeader == "" || len(authHeader) < 7 || "Bearer " != authHeader[:7] {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid authorization header"})
	}

	tokenStr := authHeader[7:]

	// парсим и проверяем jwt token
	keyFunc := func(parsedToken *jwt.Token) (interface{}, error) {
		if _, ok := parsedToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signing")
		}
		return jwtSecret, nil
	}
	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid access token"})
	}

	// добавляем данные об пользователе в контекст для следующего обработчика
	user := getUserFromPayload(token)
	
	c.Locals("user", user)

	return c.Next()
}

// TODO: добавить middleware для создания устройств
func roleMiddleware(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Проверяем заголовок c токеном
		authHeader := c.Get("Authorization")
		if authHeader == "" || len(authHeader) < 7 || "Bearer " != authHeader[:7] {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid authorization header"})
		}
	
		tokenStr := authHeader[7:]
	
		// парсим и проверяем jwt token
		keyFunc := func(parsedToken *jwt.Token) (interface{}, error) {
			if _, ok := parsedToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signing")
			}
			return jwtSecret, nil
		}
		token, err := jwt.Parse(tokenStr, keyFunc)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid access token"})
		}
	
		// получаем данные пользователя
		user := getUserFromPayload(token)

		// проверяем роль
		if user.Role != role {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "not access"})
		}
		
		// передаём данные об пользователе в следующий обработчик
		c.Locals("user", user)
	
		return c.Next()
	}
}

func authHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(User)

	token, err := genJwtToken(user.Id, user.Email, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't check token"})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

// additional

func getAllHandler(c *fiber.Ctx) error {
	users, err := userSt.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't get users"})
	}

	return c.JSON(users)
}

// func deleteHandler(c *fiber.Ctx) error {
// 	// получаем email
// 	email := c.Locals("email").(string)

// 	// проверяем, существует ли пользователь
// 	_, exists := userSt.GetUserByEmail(email)
// 	if exists != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user with provided email doesn't exists"})
// 	}

// 	// получаем id пользователя
// 	id, err := c.ParamsInt("id", 1)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
// 	}

// 	// проверяем, себя ли мы удаляем или нет
// 	// ...

// 	// удаляем пользователя
// 	err = userSt.DeleteUser(id)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't get users"})
// 	}
	
// 	return c.JSON(fiber.Map{"message": "successful deleted"})
// }

// func updateHandler(c *fiber.Ctx) error {
// 	// разбираем запрос
// 	type Request struct {
// 		Email string `json:"email"`
// 		Password string `json:"password"`
// 		Role string `json:"role"`
// 	}
	
// 	objRequest := new(Request)
// 	if err := c.BodyParser(objRequest); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input" + err.Error()})
// 	}

// 	// получаем email
// 	email := c.Locals("email").(string)

// 	// проверяем, существует ли пользователь
// 	user, exists := userSt.GetUserByEmail(email)
// 	if exists != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user with provided email doesn't exists"})
// 	}

// 	// получаем id пользователя
// 	_, err := c.ParamsInt("id", 1)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
// 	}

// 	// проверяем, себя ли мы обновляем
// 	// ...

// 	// нужно ли обновить пароль
// 	// если да, то хэшируем новый пароль
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(objRequest.Password)); err != nil {
// 		encryptedPasswrod, err := bcrypt.GenerateFromPassword([]byte(objRequest.Password), bcrypt.DefaultCost)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't update user"})
// 		}
// 		user.Password = string(encryptedPasswrod)
// 	}

// 	// генерируем новый jwt token
// 	tokenStr, err := genJwtToken(objRequest.Email)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't update user"})
// 	}

// 	// обновляем данные для пользователя в хранилище
// 	user.Email = objRequest.Email
// 	user.Password = objRequest.Password
// 	user.Role = objRequest.Role

// 	err = userSt.UpdateUser(user)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't update users"})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenStr})
// }

// helpers

func genJwtToken(id int, email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"email": email,
		"role": role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString(jwtSecret)
	return tokenStr, err
}

func getUserFromPayload(tokenJwt *jwt.Token) User {
	claims := tokenJwt.Claims.(jwt.MapClaims)

	return User{
		Id: int(claims["id"].(float64)),
		Email: claims["email"].(string),
		Role: claims["role"].(string),
	}
}