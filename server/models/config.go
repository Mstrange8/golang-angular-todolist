package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// DB variable connection for mongoDB
	DB      *mongo.Client
	cfgFile string
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
    db, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://mstrange8:18Kiwi18@cluster0.kyab7.mongodb.net/birdpedia?retryWrites=true&w=majority"))
    if err != nil {
        panic(err)
    }	
	DB = db
}