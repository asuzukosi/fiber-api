package handlers

import (
	"github/asuzukosi/fiber-api/databases"
	"github/asuzukosi/fiber-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	Id        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"createdAt"`
}

func CreateResponseOrder(order models.Order) Order {
	return Order{
		Id:        order.Id,
		User:      CreateResponseUser(order.User),
		Product:   CreateResponseProduct(order.Product),
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	user := models.User{}
	databases.Database.Db.Find(&user, "id=?", order.UserRefer)
	product := models.Product{}
	databases.Database.Db.Find(&product, "id=?", order.ProductRefer)
	order.User = user
	order.Product = product
	databases.Database.Db.Create(&order)
	return c.Status(201).JSON(CreateResponseOrder(order))
}

func GetOrder(c *fiber.Ctx) error {
	var order models.Order

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(200).JSON(err.Error())
	}
	databases.Database.Db.Find(&order, "id=?", id)
	return c.Status(200).JSON(CreateResponseOrder(order))
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	databases.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		responseOrders = append(responseOrders, CreateResponseOrder(order))
	}
	return c.Status(200).JSON(responseOrders)
}
