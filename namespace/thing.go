package namespace

import (
	"context"
	"errors"
	"github.com/appellative-ai/core/std"
	"github.com/appellative-ai/postgres/request"
	"net/url"
)

const (
	nameName       = "name"
	cNameName      = "cname"
	authorName     = "author"
	thing1Name     = "thing1"
	thing2Name     = "thing2"
	kindName       = "kind"
	collectiveName = "collective"
	domainName     = "domain"
	pathName       = "path"

	thingRequestSql = "CALL dbo.InsertThing($1,$2,$3,$4,$5,$6,$7)"
)

func thingRequest(ctx context.Context, requester *request.Interface, values url.Values) (request.Result, error) {
	if values == nil {
		return request.Result{}, errors.New("query values are nil")
	}
	name, args, err := createThingArgs(values)
	if err != nil {
		return request.Result{}, err
	}
	return requester.Execute(ctx, name, thingRequestSql, args...)
}

func createThingArgs(values url.Values) (string, []any, error) {
	name := values.Get(nameName)
	if name == "" {
		return "", nil, errors.New("name is empty")
	}
	author := values.Get(authorName)
	if author == "" {
		return "", nil, errors.New("author is empty")
	}
	n := std.NewName(name)
	var args []any

	args = append(args, name)
	cname := values.Get(cNameName)
	if cname != "" {
		args = append(args, cname)
	}
	args = append(args, author)
	args = append(args, n.Collective)
	args = append(args, n.Domain)
	args = append(args, n.Kind)
	args = append(args, n.Path)
	return name, args, nil
}
