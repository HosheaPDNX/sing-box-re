package log

import (
	"context"
	"math/rand"
	"time"
)

type idKey struct{}

type parentIDKey struct{}

type ID struct {
	ID        uint32
	CreatedAt time.Time
}

func ContextWithNewID(ctx context.Context) context.Context {
	return ContextWithID(ctx, ID{
		ID:        rand.Uint32(),
		CreatedAt: time.Now(),
	})
}

func ContextWithNewChildID(ctx context.Context) context.Context {
	parentID, parentLoaded := IDFromContext(ctx)
	ctx = ContextWithNewID(ctx)
	if !parentLoaded {
		return ctx
	}
	return context.WithValue(ctx, (*parentIDKey)(nil), parentID)
}

func ContextWithID(ctx context.Context, id ID) context.Context {
	return context.WithValue(ctx, (*idKey)(nil), id)
}

func IDFromContext(ctx context.Context) (ID, bool) {
	id, loaded := ctx.Value((*idKey)(nil)).(ID)
	return id, loaded
}

func ParentIDFromContext(ctx context.Context) (ID, bool) {
	id, loaded := ctx.Value((*parentIDKey)(nil)).(ID)
	return id, loaded
}
