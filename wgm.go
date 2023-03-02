package wgm

import (
	"context"
	"github.com/gookit/slog"
	"github.com/qiniu/qmgo"
	"time"
)

var instance *wgm

type wgm struct {
	client *qmgo.Client
	dbName string
}

func InitWgm(connectionUri string, databaseName string) error {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri:      connectionUri,
		Database: databaseName,
	})
	if err != nil {
		return err
	}
	instance = &wgm{
		client: client,
		dbName: databaseName,
	}
	return nil
}

func (w *wgm) newCtxWithTimeout(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ctx
}

func (w *wgm) newCtx() context.Context {
	return w.newCtxWithTimeout(10 * time.Second)
}

func CloseAll() {
	if instance == nil {
		slog.Fatal("must initialize WGM first, by calling InitWgm() method")
	}

	if err := instance.client.Close(instance.Ctx()); err != nil {
		slog.Error(err)
	}
}

func Ping() error {
	if instance == nil {
		slog.Fatal("must initialize WGM first, by calling InitWgm() method")
	}
	return instance.client.Ping(10)
}
