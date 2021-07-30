package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetCollection(collection string) *mongo.Collection {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://blas:9HzjhrnrDjXX0C5z@cluster0.hbnsh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	if err != nil {
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("connected to mongodb")
	return client.Database("users").Collection(collection)
}
