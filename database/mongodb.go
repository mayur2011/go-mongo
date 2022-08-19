package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const MongoURI = "mongodb://readonly:readonly@cluster01-shard-00-00.wvrzd.mongodb.net:27017,cluster01-shard-00-01.wvrzd.mongodb.net:27017,cluster01-shard-00-02.wvrzd.mongodb.net:27017/demo_db?replicaSet=demo_rs&authSource=admin&readPreference=secondaryPreferred&maxIdleTimeMS=120000&maxPoolSize=100&socketTimeoutMS=5000"

func ConnectDB(ctx context.Context) *mongo.Database {

	connection := options.Client().ApplyURI(MongoURI)
	client, err := mongo.Connect(ctx, connection)
	if err != nil {
		panic(err)
	}
	return client.Database("goss_prod")
}

func NewMongoClient(ctx context.Context) (*mongo.Client, error) {

	//step-01: mongo uri and setup mongo client with mongo.NewClient
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	if err != nil {
		log.Println(err)
	}

	//step-02: mongo client to connect with context timeout setting
	if err = mongoClient.Connect(ctx); err != nil {
		log.Println(err)
	}

	//step-03: mongo client - perform ping
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	fmt.Println("mongoClient Ping")

	return mongoClient, err
}

func GetDatabaseName(ctx context.Context, mongoClient *mongo.Client) []string {
	//step-04: mongo client to list database names
	db, err := mongoClient.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	return db
}
