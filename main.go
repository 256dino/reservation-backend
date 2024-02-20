package main

import (
	"context"
	"flag"
	"github.com/256dino/reservation-backend/api"
	"github.com/256dino/reservation-backend/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const dburi = "mongodb://localhost:27017"

//var config = fiber.Config{
//	ErrorHandler: func(c *fiber.Ctx, err error) error {
//		return c.JSON(map[string]string{"error": err.Error()})
//	},
//}

func main() {
	listenAddr := flag.String("listenAddr", ":9000", "port address")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user/", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)

	app.Listen(*listenAddr) // Should be last in main()
}
