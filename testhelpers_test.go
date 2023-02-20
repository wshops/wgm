package wgm_test

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

func InsertStu(stu *Student) primitive.ObjectID {
	result, err := wgm.Insert(stu)
	if err != nil {
		panic(err)
	}
	return result.InsertedID.(primitive.ObjectID)
}

func DelStu(stu *Student) {
	err := wgm.Delete(stu)
	if err != nil {
		panic(err)
	}
}
