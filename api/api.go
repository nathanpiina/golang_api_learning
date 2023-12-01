package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathanpiina/golang_api_learning/person"
)

func Main() {
	app := fiber.New()

	app.Post("/people", func(c *fiber.Ctx) error {
		nickname := c.FormValue("nickname")
		name := c.FormValue("name")
		birth := c.FormValue("birth")
		stack := c.FormValue("stack")
	
		err := person.AddPeople(nickname, name, birth, stack)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	
		return c.SendString("Usuário adicionado com sucesso")
	})
	
	app.Get("/people/:nickname", func(c *fiber.Ctx) error {
		nickname := c.Params("nickname")
		result, err := person.SearchPeople(nickname)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
	
		return c.SendString(result)
	})

	app.Get("/count", func(c *fiber.Ctx) error {
		result, err := person.CountRowsInTable()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	
		return c.SendString(result)
	})

	person.Main()

	app.Listen(":3000")
}