package handlers

import (
	"github/asuzukosi/fiber-api/databases"
	"github/asuzukosi/fiber-api/models"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	// This is the user schema
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func CreateResponseUser(user models.User) User {
	return User{Id: user.Id, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	databases.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(201).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	databases.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}

	return c.JSON(responseUsers)
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	user := models.User{}
	databases.Database.Db.Find(&user, "id = ?", id)
	if user.Id == 0 {
		return c.Status(400).SendString("User does not exist")
	}
	responseUser := CreateResponseUser(user)

	return c.JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	databases.Database.Db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	databases.Database.Db.Find(&user, "id = ?", id)

	return c.JSON(CreateResponseUser(user))
}

func DeleteUser(c *fiber.Ctx) error {
	user := models.User{}

	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	databases.Database.Db.Delete(&user, "id = ?", id)

	return c.SendString("User deleted successfully")
}
