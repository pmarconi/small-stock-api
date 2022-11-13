package handlers

import (
	"stock-manager-api/database"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "",
		Data:    []database.Product{},
	})
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	message := "Hello " + id
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: message,
		Data:    []database.Product{},
	})
}
