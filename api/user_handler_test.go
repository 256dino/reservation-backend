package api

import (
	"context"
	"github.com/256dino/reservation-backend/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

const testmMongoURI = "mongodb://localhost:27017"

type testdb struct {
	db.UserStore
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testmMongoURI))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client),
	}
}

func TestUserHandler_HandlePostUser(t *testing.T) {
	testdb := setup(t)
	defer testdb.teardown(t)
	t.Fail()
}

func (tbd *testdb) teardown(t *testing.T) {
	if err := tbd.UserStore.Drop(context.Background()); err != nil {
		log.Fatal(err)
	}
}
