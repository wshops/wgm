package wgm_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wshops/wgm"
	"github.com/wshops/wgm/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestGetCollection(t *testing.T) {
	internal.SetupDefaultConnection()
	doc := internal.NewDoc("Alice", 12)
	collection := wgm.Col(doc.ColName())

	require.NotNil(t, collection)
}

func TestGetCtx(t *testing.T) {
	internal.SetupDefaultConnection()
	ctx := wgm.Ctx()
	ctx.Deadline()

	_, ok := ctx.Deadline()
	require.True(t, ok, "context should having deadline.")
}

func TestFindWithPage(t *testing.T) {
	internal.SetupDefaultConnection()
	doc1 := internal.NewDoc("Alice", 12)
	doc2 := internal.NewDoc("Candy", 24)
	internal.InsertDoc(doc1)
	internal.InsertDoc(doc2)
	defer internal.DelDoc(doc1)
	defer internal.DelDoc(doc2)

	var doc []*internal.Doc
	wgm.FindWithPage(&internal.Doc{}, nil, &doc, 5, 1)

	var DBDocs = []*internal.Doc{doc1, doc2}
	for i, dbDoc := range DBDocs {
		require.Equal(t, doc[i].Age, dbDoc.Age)
	}
}

func TestFindOne(t *testing.T) {
	internal.SetupDefaultConnection()
	DBdoc := internal.NewDoc("Alice", 12)
	objectID := internal.InsertDoc(DBdoc)
	defer internal.DelDoc(DBdoc)

	var doc internal.Doc
	wgm.FindOne(&doc, nil)
	require.Equal(t, doc.Id, objectID)
}

func TestFindById(t *testing.T) {
	internal.SetupDefaultConnection()
	DBdoc := internal.NewDoc("Alice", 12)
	objectID := internal.InsertDoc(DBdoc)
	defer internal.DelDoc(DBdoc)

	doc := &internal.Doc{}
	doc.Id = objectID
	_, err := wgm.FindById(doc.ColName(), doc.GetId(), doc)
	require.Nil(t, err)

	require.Equal(t, doc.Name, DBdoc.Name)
	require.Equal(t, doc.Age, DBdoc.Age)
}

func TestInsert(t *testing.T) {
	internal.SetupDefaultConnection()
	doc := internal.NewDoc("Alice", 12)
	defer internal.DelDoc(doc)

	result, err := wgm.Insert(doc)
	require.Nil(t, err)

	DBdoc := &internal.Doc{}
	wgm.FindOne(DBdoc, nil)
	require.Equal(t, result.InsertedID.(primitive.ObjectID), DBdoc.Id)
}

func TestUpdate(t *testing.T) {
	internal.SetupDefaultConnection()
	DBdoc := internal.NewDoc("Alice", 12)
	ObjectID := internal.InsertDoc(DBdoc)
	defer internal.DelDoc(DBdoc)

	var doc = &internal.Doc{
		Name: "Bob",
		Age:  99,
	}
	doc.Id = ObjectID
	err := wgm.Update(doc)
	require.Nil(t, err)

	DBdoc = &internal.Doc{}
	wgm.FindOne(DBdoc, nil)
	require.Equal(t, DBdoc.Age, doc.Age)
	require.Equal(t, DBdoc.Name, doc.Name)
}

func TestDelete(t *testing.T) {
	internal.SetupDefaultConnection()
	DBdoc := internal.NewDoc("Alice", 12)
	ObjectID := internal.InsertDoc(DBdoc)

	DBdoc.Id = ObjectID
	err := wgm.Delete(DBdoc)
	require.Nil(t, err)

	result := wgm.FindOne(DBdoc, bson.M{"_id": ObjectID})
	require.False(t, result)
}
