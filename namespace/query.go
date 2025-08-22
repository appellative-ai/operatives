package namespace

import (
	"bytes"
	"context"
	"github.com/appellative-ai/operatives/template"
	"github.com/appellative-ai/postgres/retrieval"
	"net/http"
)

// queryRetrieval - applies to current and all linked collectives
func queryRetrieval(ctx context.Context, retriever *retrieval.Interface, processor template.Agent, r *http.Request) (*bytes.Buffer, error) {
	name := ""
	res, err := processor.Build(name, nil)
	if err != nil {
		return nil, err
	}
	return retriever.Marshal(ctx, name, res.Sql, res.Args)
}
