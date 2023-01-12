package internal

import "github.com/wshops/wgm"

type Doc struct {
	wgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name"`
	Age              int    `bson:"age"`
}

func (d *Doc) ColName() string {
	return "Docs"
}

func NewDoc(name string, age int) *Doc {
	return &Doc{Name: name, Age: age}
}
