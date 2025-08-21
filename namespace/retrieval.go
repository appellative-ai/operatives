package namespace

import (
	"bytes"
	"context"
	"github.com/appellative-ai/center/template"
	"github.com/appellative-ai/postgres/retrieval"
	"net/http"
)

type tagRetrieval struct {
	Name string `json:"name"`
	Args []arg  `json:"args"`
}

func retrievalRequest(ctx context.Context, retriever *retrieval.Interface, processor template.Agent, r *http.Request) (*bytes.Buffer, error) {
	name := ""
	res, err := processor.Build(name, nil)
	if err != nil {
		return nil, err
	}
	return retriever.Marshal(ctx, name, res.Sql, res.Args)
}
