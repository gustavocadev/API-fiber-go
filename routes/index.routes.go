package routes

import (
	// "goFiber/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

// users
var users = []*User{
	{
		Id:   xid.New().String(),
		Name: "Agustin",
		Age:  29,
	},
}

func UseRoute(router fiber.Router) {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": users,
		})
	})

	router.Post("/", func(c *fiber.Ctx) error {
		type Request struct {
			Name string
			Age  uint8
		}

		var body Request

		err := c.BodyParser(&body)

		if err != nil {

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Hay un error",
			})
		}

		if len(body.Name) == 0 {
			return c.JSON(fiber.Map{
				"error": "Necesitas agregar contenido al name",
			})
		}

		newUser := &User{
			Id:   xid.New().String(),
			Name: body.Name,
			Age:  body.Age,
		}

		users = append(users, newUser)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": users,
		})
	})

	router.Put("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		type Request struct {
			Name string
			Age  uint8
		}
		var body Request
		c.BodyParser(&body)

		for _, user := range users {
			if user.Id == id {
				user.Name = body.Name
				user.Age = body.Age
			}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": users,
		})
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for idx, user := range users {
			if user.Id == id {
				users = append(users[:idx], users[idx+1:]...)
			}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": users,
		})
	})
}
