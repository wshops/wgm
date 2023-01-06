package wgm

import (
    "context"
    "errors"

    "github.com/gookit/slog"
)

type updater struct {
    collectionModel any
    ctx             context.Context
    hasResult       bool
}

// Updater creates a new updater instance, can be used for chain call
func Updater(m any) *updater {
    if m == nil {
        slog.Error("must provide model to updater")
        return nil
    }
    return &updater{collectionModel: m, ctx: Ctx()}
}

// Find based on the model provided, execute find one in database
// return *updater the updater obj for chain call
// return bool does document exist in the database
func (u *updater) Find() (*updater, bool) {
    if u.collectionModel == nil {
        slog.Error("must provide model to updater")
        return u, false
    }
    hasResult, err := FindById(u.collectionModel.(IDefaultModel).ColName(), u.collectionModel.(IDefaultModel).GetId(), u.collectionModel)
    if !hasResult {
        return nil, false
    }
    if err != nil {
        slog.Error(err)
        return nil, false
    }
    u.hasResult = true
    return u, true
}

// Update updates the model modified to the database
func (u *updater) Update(filter ...map[string]any) error {
    if !u.hasResult {
        return errors.New("document does not exist")
    }
    if u.collectionModel == nil {
        return errors.New("must provide model to updater")
    }
    return Update(u.collectionModel.(IDefaultModel), filter...)
}
