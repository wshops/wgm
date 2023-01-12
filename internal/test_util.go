package internal

import (
	"github.com/wshops/wgm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetupDefaultConnection() {
	err := wgm.InitWgm("mongodb://localhost:27017/", "wshop_test")
	if err != nil {
		panic(err)
	}
}

func InsertDoc(doc *Doc) primitive.ObjectID {
	result, err := wgm.Insert(doc)
	if err != nil {
		panic(err)
	}
	return result.InsertedID.(primitive.ObjectID)
}

func DelDoc(doc *Doc) {
	err := wgm.Delete(doc)
	if err != nil {
		panic(err)
	}
}
