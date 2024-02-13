package api

import (
	"github.com/256dino/reservation-backend/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "joe",
		Lastname:  "bob",
	}
	return c.JSON(u)
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("Joe")
}
