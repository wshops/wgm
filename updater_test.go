package wgm_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wshops/wgm"
	"github.com/wshops/wgm/internal"
	"testing"
)

func TestGetUpdater(t *testing.T) {
	doc := internal.Doc{}
	u := wgm.Updater(doc)
	require.NotNil(t, u)

}

func TestUpdaterFind(t *testing.T) {
	DBdoc := internal.NewDoc("Alice", 12)
	internal.SetupDefaultConnection()
	objectID := internal.InsertDoc(DBdoc)
	defer internal.DelDoc(DBdoc)

	var doc internal.Doc
	doc.Id = objectID
	u := wgm.Updater(doc)
	u.Find()

	require.Equal(t, doc.Age, DBdoc.Age)
	require.Equal(t, doc.Name, DBdoc.Name)
}

func TestUpdater_Update(t *testing.T) {
	DBdoc := internal.NewDoc("Alice", 12)
	internal.SetupDefaultConnection()
	objectID := internal.InsertDoc(DBdoc)
	defer internal.DelDoc(DBdoc)

	var doc internal.Doc
	doc.Id = objectID
	u := wgm.Updater(doc)
	u.Find()
	doc.Name = "dada"
	doc.Age = 99
	err := u.Update()
	require.Nil(t, err)

	DBdoc = &internal.Doc{}
	wgm.FindOne(DBdoc, nil)
	require.Equal(t, DBdoc.Age, 99)
	require.Equal(t, DBdoc.Name, "dada")
}
