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

	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	if err := app.Listen(*listenAddr); err != nil {
		log.Fatal(err)
	}
}
