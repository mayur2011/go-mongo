package main

import (
	"context"
	"fmt"
	"go-mongodb-rmq/database"
	"log"
	"time"
)

func main() {

	//setting context with timeout of 10 seconds
	ctx, stop := context.WithTimeout(context.Background(), 10*time.Second)
	defer stop()

	mongoClient, err := database.NewMongoClient(ctx)
	if err != nil {
		log.Println("error with mongo client", err)
	}
	defer mongoClient.Disconnect(ctx)

	fmt.Println(database.GetDatabaseName(ctx, mongoClient))
}

/*
fixing error by adding
go get go.mongodb.org/mongo-driver/mongo/options@v1.10.1
go get go.mongodb.org/mongo-driver/mongo@v1.10.1
*/
