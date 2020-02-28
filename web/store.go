package web

import (
	"bytes"
	"context"

	"github.com/CliffLin/go/backend"
	"github.com/CliffLin/go/internal"
)

// Store defines the link store interface.
type Store interface {
	GetLink(ctx context.Context, key string) (*internal.Route, error)
	GetLinks(ctx context.Context, str string) (map[string]internal.Route, error)
	UpdateLink(ctx context.Context, key string, route *internal.Route) error
	DeleteLink(ctx context.Context, key string) error
	NextID(ctx context.Context) (uint64, error)
}

// NewStore returns a new store with kv and enc.
func NewStore(backend backend.Backend) Store {
	return &store{
		backend: backend,
	}
}

type store struct {
	backend backend.Backend
}

func (s *store) GetLink(ctx context.Context, key string) (route *internal.Route, err error) {
	b, err := s.backend.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	rt := &internal.Route{}
	if err = rt.Read(bytes.NewBuffer(b)); err != nil {
		return nil, err
	}
	return rt, err
}

func (s *store) GetLinks(ctx context.Context, start string) (routes map[string]internal.Route, err error) {
	return s.backend.List(ctx, start)
}

func (s *store) UpdateLink(ctx context.Context, key string, route *internal.Route) (err error) {
	var buf bytes.Buffer
	if err = route.Write(&buf); err != nil {
		return
	}
	return s.backend.Put(ctx, key, buf.Bytes())
}

func (s *store) DeleteLink(ctx context.Context, key string) (err error) {
	return s.backend.Del(ctx, key)
}

func (s *store) NextID(ctx context.Context) (uint64, error) {
	return s.backend.NextID(ctx)
}
