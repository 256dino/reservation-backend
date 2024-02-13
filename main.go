package main

import (
	"flag"
	"github.com/256dino/reservation-backend/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":7000", "port address")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)
	apiv1.Get("/user", api.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "hey"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "joe18318"})
}
