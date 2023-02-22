package userrepords

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MongoDBRepository struct {
		doc *mongo.Collection
	}
)

type DocumentUser struct {
	ID   primitive.ObjectID     `json:"id" bson:"_id"`
	Item string                 `json:"item" bson:"item"`
	Qty  int32                  `json:"qty" bson:"qty"`
	Tags []string               `json:"tags" bson:"tags"`
	Size map[string]interface{} `json:"size" bson:"size"`
}

var (
	ctx                 = context.TODO()
	mongoCollectionName = "capital"
)

func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		doc: db.Collection(mongoCollectionName),
	}
}

func (repo *MongoDBRepository) List() []DocumentUser {
	var docUser []DocumentUser

	list, err := repo.doc.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}

	defer list.Close(ctx)

	for list.Next(ctx) {
		user := DocumentUser{}
		err = list.Decode(&user)
		if err != nil {
			fmt.Printf("an error ocurred when parse data from cursor : %v", err)
			panic(err)
		}
		docUser = append(docUser, user)
	}

	return docUser
}
