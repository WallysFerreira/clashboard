package db

import (
  "os"
  "log"
  "context"

  "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
  godotenv.Load()

  uri := os.Getenv("MONGO_URI")
  if uri == "" {
    log.Fatal("No MONGO_URI env var set")
  }

  client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
  if err != nil {
    panic(err)
  }

  return client
}
