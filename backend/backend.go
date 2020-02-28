package backend

import (
	"context"

	"github.com/CliffLin/go/internal"
)

type Backend interface {
	Close() error
	Get(ctx context.Context, id string) ([]byte, error)
	Put(ctx context.Context, key string, buf []byte) error
	Del(ctx context.Context, id string) error
	GetAll(ctx context.Context) (map[string]internal.Route, error)
	List(ctx context.Context, start string) (map[string]internal.Route, error)
	NextID(ctx context.Context) (uint64, error)
}
