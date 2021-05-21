package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func newMongoDatabase() (*mongo.Database, error) {
	mongoUrl := "mongodb://localhost/" //os.Getenv("MONGO_URL")
	mongodb := "go-store"              //os.Getenv("MONGO_DB")
	mongoTimeout := 30                 //strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	database := client.Database(mongodb)
	return database, nil
}
