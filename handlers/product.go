package handlers

import (
	"net/http"
	"stock-manager-api/database"
	"stock-manager-api/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	product := new(database.Product)

	if err := c.BodyParser(product); err != nil {
		utils.Error().Println(err.Error())
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "Error",
			Data:    nil,
		})
	}

	database.AddProduct(product)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Product added",
		Data:    []database.Product{},
	})
}

func ReadProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Why use int64?
	id64, err := strconv.ParseUint(id, 0, 64)

	if err != nil {
		utils.Error().Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "Error",
			Data:    nil,
		})
	}

	productId := uint(id64)
	product, err := database.GetProductById(&productId)
	if err != nil {
		utils.Error().Println(err.Error())
		return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
			Success: false,
			Message: "Error",
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(ResponseHTTP{
		Success: true,
		Message: "",
		Data:    product,
	})
}

func ReadProducts(c *fiber.Ctx) error {
	products, err := database.GetProducts()
	if err != nil {
		utils.Error().Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "Error",
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(ResponseHTTP{
		Success: true,
		Message: "",
		Data:    products,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	// TODO:
	return c.Status(http.StatusOK).JSON(ResponseHTTP{
		Success: true,
		Message: "",
		Data:    []database.Product{},
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	// TODO:
	id := c.Params("id")
	return c.Status(http.StatusOK).JSON(ResponseHTTP{
		Success: true,
		Message: "" + id,
		Data:    []database.Product{},
	})
}
