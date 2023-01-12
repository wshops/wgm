package wgm_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wshops/wgm"
	"testing"
)

func TestInitWgm(t *testing.T) {
	err := wgm.InitWgm("mongodb://localhost:27017/", "wshop_test")
	require.Nil(t, err)
}

func TestSetupWrongConnection(t *testing.T) {
	err := wgm.InitWgm("wrong://localhost:27017/", "wshop_test")
	require.NotNil(t, err)
}

func TestPing(t *testing.T) {
	SetupDefaultConnection()
	err := wgm.Ping()
	require.Nil(t, err)
}
