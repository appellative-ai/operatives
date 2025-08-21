package namespacetest

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/appellative-ai/core/std"
	"github.com/appellative-ai/postgres/request"
	"github.com/appellative-ai/postgres/retrieval"
)

var (
	store = newStorage()
)

type storage struct {
	cache *std.MapT[string, any]
}

func newStorage() *storage {
	s := new(storage)
	s.cache = std.NewSyncMap[string, any]()
	return s
}

func (s *storage) Marshal(ctx context.Context, name, sql string, args ...any) (bytes.Buffer, error) {
	t, _ := s.cache.Load(name)
	buf, err := json.Marshal(t)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return *bytes.NewBuffer(buf), err
}

func (s *storage) Scan(ctx context.Context, fn retrieval.ScanFunc, name, sql string, args ...any) error {
	return nil
}

func (s *storage) Execute(ctx context.Context, name, sql string, args ...any) (request.Result, error) {
	return request.Result{}, nil
}
