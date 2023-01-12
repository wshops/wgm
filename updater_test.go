package wgm_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wshops/wgm"
	"testing"
)

func TestGetUpdater(t *testing.T) {
	doc := Doc{}
	u := wgm.Updater(doc)
	require.NotNil(t, u)

}

func TestUpdaterFind(t *testing.T) {
	DBdoc := NewDoc("Alice", 12)
	SetupDefaultConnection()
	objectID := InsertDoc(DBdoc)
	defer DelDoc(DBdoc)

	doc := &Doc{}
	doc.Id = objectID
	u, hasResult := wgm.Updater(doc).Find()

	require.True(t, hasResult)
	require.NotNil(t, u)
	require.Equal(t, doc.Age, DBdoc.Age)
	require.Equal(t, doc.Name, DBdoc.Name)
}

func TestUpdater_Update(t *testing.T) {
	DBdoc := NewDoc("Alice", 12)
	SetupDefaultConnection()
	objectID := InsertDoc(DBdoc)
	defer DelDoc(DBdoc)

	doc := &Doc{}
	doc.Id = objectID
	u, _ := wgm.Updater(doc).Find()
	doc.Name = "dada"
	doc.Age = 99
	err := u.Update()
	require.Nil(t, err)

	DBdoc = &Doc{}
	wgm.FindOne(DBdoc, nil)
	require.Equal(t, DBdoc.Age, 99)
	require.Equal(t, DBdoc.Name, "dada")
}
