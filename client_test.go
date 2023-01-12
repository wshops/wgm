package wgm_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wshops/wgm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestGetCollection(t *testing.T) {
	SetupDefaultConnection()
	doc := NewDoc("Alice", 12)
	collection := wgm.Col(doc.ColName())

	require.NotNil(t, collection)
}

func TestGetCtx(t *testing.T) {
	SetupDefaultConnection()
	ctx := wgm.Ctx()
	ctx.Deadline()

	_, ok := ctx.Deadline()
	require.True(t, ok, "context should having deadline.")
}

func TestFindWithPage(t *testing.T) {
	SetupDefaultConnection()
	doc1 := NewDoc("Alice", 12)
	doc2 := NewDoc("Candy", 24)
	InsertDoc(doc1)
	InsertDoc(doc2)
	defer DelDoc(doc1)
	defer DelDoc(doc2)

	var doc []*Doc
	wgm.FindWithPage(&Doc{}, nil, &doc, 5, 1)

	var DBDocs = []*Doc{doc1, doc2}
	for i, dbDoc := range DBDocs {
		require.Equal(t, doc[i].Age, dbDoc.Age)
		require.Equal(t, doc[i].Name, dbDoc.Name)
	}
}

func BenchmarkFindWithPage(b *testing.B) {
	SetupDefaultConnection()
	doc1 := NewDoc("Alice", 12)
	doc2 := NewDoc("Candy", 24)
	InsertDoc(doc1)
	InsertDoc(doc2)
	defer DelDoc(doc1)
	defer DelDoc(doc2)

	var doc []*Doc
	wgm.FindWithPage(&Doc{}, nil, &doc, 5, 1)

	var DBDocs = []*Doc{doc1, doc2}
	for i, dbDoc := range DBDocs {
		require.Equal(b, doc[i].Age, dbDoc.Age)
		require.Equal(b, doc[i].Name, dbDoc.Name)
	}
}

func TestFindOne(t *testing.T) {
	SetupDefaultConnection()
	DBdoc := NewDoc("Alice", 12)
	objectID := InsertDoc(DBdoc)
	defer DelDoc(DBdoc)

	var doc Doc
	wgm.FindOne(&doc, nil)
	require.Equal(t, doc.Id, objectID)
}

func BenchmarkFindOne(b *testing.B) {
	SetupDefaultConnection()
	DBdoc := NewDoc("Alice", 12)
	objectID := InsertDoc(DBdoc)
	defer DelDoc(DBdoc)

	var doc Doc
	wgm.FindOne(&doc, nil)
	require.Equal(b, doc.Id, objectID)
}

func TestFindById(t *testing.T) {
	SetupDefaultConnection()
	DBdoc := NewDoc("Alice", 12)
	objectID := InsertDoc(DBdoc)
	defer DelDoc(DBdoc)

	doc := &Doc{}
	doc.Id = objectID
	_, err := wgm.FindById(doc.ColName(), doc.GetId(), doc)
	require.Nil(t, err)

	require.Equal(t, doc.Name, DBdoc.Name)
	require.Equal(t, doc.Age, DBdoc.Age)
}

func BenchmarkFindById(b *testing.B) {
	SetupDefaultConnection()
	DBdoc := NewDoc("Alice", 12)
	objectID := InsertDoc(DBdoc)
	defer DelDoc(DBdoc)

	doc := &Doc{}
	doc.Id = objectID
	_, err := wgm.FindById(doc.ColName(), doc.GetId(), doc)
	require.Nil(b, err)

	require.Equal(b, doc.Name, DBdoc.Name)
	require.Equal(b, doc.Age, DBdoc.Age)
}

func TestInsert(t *testing.T) {
	SetupDefaultConnection()
	doc := NewDoc("Alice", 12)
	defer DelDoc(doc)

	result, err := wgm.Insert(doc)
	require.Nil(t, err)

	DBdoc := &Doc{}
	wgm.FindOne(DBdoc, nil)
	require.Equal(t, result.InsertedID.(primitive.ObjectID), DBdoc.Id)
}

func BenchmarkInsert(b *testing.B) {
	SetupDefaultConnection()
	doc := NewDoc("Alice", 12)
	defer DelDoc(doc)

	result, err := wgm.Insert(doc)
	require.Nil(b, err)

	DBdoc := &Doc{}
	wgm.FindOne(DBdoc, nil)
	require.Equal(b, result.InsertedID.(primitive.ObjectID), DBdoc.Id)
}

func TestUpdate(t *testing.T) {
	SetupDefaultConnection()
	DBdoc := NewDoc("Alice", 12)
	ObjectID := InsertDoc(DBdoc)
	defer DelDoc(DBdoc)

	var doc = &Doc{
		Name: "Bob",
		Age:  99,
	}
	doc.Id = ObjectID
	err := wgm.Update(doc)
	require.Nil(t, err)

	DBdoc = &Doc{}
	wgm.FindOne(DBdoc, nil)
	require.Equal(t, DBdoc.Age, doc.Age)
	require.Equal(t, DBdoc.Name, doc.Name)
}

func BenchmarkUpdate(b *testing.B) {
	SetupDefaultConnection()
	DBdoc := NewDoc("Alice", 12)
	ObjectID := InsertDoc(DBdoc)
	defer DelDoc(DBdoc)

	var doc = &Doc{
		Name: "Bob",
		Age:  99,
	}
	doc.Id = ObjectID
	err := wgm.Update(doc)
	require.Nil(b, err)

	DBdoc = &Doc{}
	wgm.FindOne(DBdoc, nil)
	require.Equal(b, DBdoc.Age, doc.Age)
	require.Equal(b, DBdoc.Name, doc.Name)
}

func TestDelete(t *testing.T) {
	SetupDefaultConnection()
	DBdoc := NewDoc("Alice", 12)
	ObjectID := InsertDoc(DBdoc)

	DBdoc.Id = ObjectID
	err := wgm.Delete(DBdoc)
	require.Nil(t, err)

	result := wgm.FindOne(DBdoc, bson.M{"_id": ObjectID})
	require.False(t, result)
}

func BenchmarkDelete(b *testing.B) {
	SetupDefaultConnection()
	DBdoc := NewDoc("Alice", 12)
	ObjectID := InsertDoc(DBdoc)

	DBdoc.Id = ObjectID
	err := wgm.Delete(DBdoc)
	require.Nil(b, err)

	result := wgm.FindOne(DBdoc, bson.M{"_id": ObjectID})
	require.False(b, result)
}
