package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type welcomeMessage struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func WelcomeHandler(c *fiber.Ctx) error {
	return c.JSON(welcomeMessage{"Hello", time.Now()})
}
