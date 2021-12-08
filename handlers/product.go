package handlers

import (
	"github/asuzukosi/fiber-api/databases"
	"github/asuzukosi/fiber-api/models"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateResponseProduct(p models.Product) Product {
	return Product{
		Id:   p.Id,
		Name: p.Name,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	product := models.Product{}

	err := c.BodyParser(&product)

	if err != nil {
		return err
	}

	databases.Database.Db.Create(&product)

	return c.Status(201).JSON(CreateResponseProduct(product))
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	databases.Database.Db.Find(&products)

	response_products := []Product{}
	for _, product := range products {
		response_products = append(response_products, CreateResponseProduct(product))
	}

	return c.Status(200).JSON(response_products)
}

func GetProduct(c *fiber.Ctx) error {
	product := models.Product{}

	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	databases.Database.Db.Find(&product, "id = ?", id)

	return c.Status(200).JSON(CreateResponseProduct(product))
}

func UpdateProduct(c *fiber.Ctx) error {
	product := models.Product{}
	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	err = c.BodyParser(&product)

	if err != nil {
		return err
	}

	databases.Database.Db.Model(&models.Product{}).Find("id = ?", id).Updates(product)
	databases.Database.Db.Find(&product, "id = ?", id)

	return c.Status(200).JSON(CreateResponseProduct(product))
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	databases.Database.Db.Delete(&models.Product{}, "id = ?", id)

	return c.Status(204).SendString("Product deleted successfully!")
}
