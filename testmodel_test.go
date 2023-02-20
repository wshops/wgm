package wgm_test

import (
	"github.com/wshops/wgm"
)

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

type Student struct {
	wgm.DefaultModel `bson:",inline"`
	Info             Info `bson:"info"`
}

func NewStudent(info Info) *Student {
	return &Student{Info: info}
}

func (s *Student) ColName() string {
	return "Students"
}

type Info struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}
