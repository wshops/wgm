package wgm_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wshops/wgm/internal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestPutId(t *testing.T) {
	doc := internal.Doc{}
	hexId := "63632c7dfc826378c8abd802"
	doc.PutId(hexId)
	hex, _ := primitive.ObjectIDFromHex(hexId)
	require.Equal(t, hex, doc.Id)
}
