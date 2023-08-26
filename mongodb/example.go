package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type Model struct {
	Id          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}

func main() {
	db, err := New(&Config{
		URL:      "mongodb://username:pwd@localhost:27017/dbname",
		Database: "dbname",
	})
	if err != nil {
		log.Println(err.Error())
		return
	}

	var result *Model
	err = db.FindOne(
		context.Background(),
		"collection",
		&result,
		WithFilter(bson.M{"name": "Quang"}),
		WithHint("_index_"),
	)

	if err != nil {
		log.Println(err.Error())
		return
	}
}
