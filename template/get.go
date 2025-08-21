package template

import (
	"encoding/json"
	"github.com/appellative-ai/core/httpx"
)

// TODO: check for not found
func (a *agentT) get(name string) (Entry, error) {
	ctx, cancel := httpx.NewContext(nil, a.timeout)
	defer cancel()

	buf, err := a.retriever.Marshal(ctx, name, defaultSql, name)
	if err != nil {
		return Entry{}, err
	}
	var t Entry
	err = json.Unmarshal(buf.Bytes(), &t)
	return t, err
}
