package config

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx     = context.Background()
	DBMongo *mongo.Database
)

func OpenMongoDBPool() {
	DBMongo = ConnectMongoDB()
}

func ConnectMongoDB() *mongo.Database {
	host := viper.GetString("MONGO_ENV")
	port := viper.GetString("MONGO_HOST")
	database := viper.GetString("MONGO_DATABASE")

	opt := options.Client()
	opt.ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))

	db, err := mongo.NewClient(opt)
	if err != nil {
		panic(err)
	}

	err = db.Connect(ctx)
	if err != nil {
		panic(err)
	}

	return db.Database(database)
}
