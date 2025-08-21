package namespace

import (
	"context"
	"errors"
	"github.com/appellative-ai/postgres/request"
	"net/url"
)

// Note: args 1 - 7 are the same as a thing, with args 8 and 9 are the thing1 and thing2
//       So a link is just a thing with the 2 additional thing links

const (
	requestLinkSql = "CALL dbo.InsertLink($1,$2,$3,$4,$5,$6,$7,$8,$9)"
)

func linkRequest(ctx context.Context, requester *request.Interface, values url.Values) (request.Result, error) {
	if values == nil {
		return request.Result{}, errors.New("query values are nil")
	}
	name, args, err := createLinkArgs(values)
	if err != nil {
		return request.Result{}, err
	}
	return requester.Execute(ctx, name, requestLinkSql, args...)
}

func createLinkArgs(values url.Values) (string, []any, error) {
	name, args, err := createThingArgs(values)
	if err != nil {
		return name, args, err
	}
	thing1 := values.Get(thing1Name)
	if thing1 == "" {
		return "", nil, errors.New("thing1 is empty")
	}
	thing2 := values.Get(thing2Name)
	if thing2 == "" {
		return "", nil, errors.New("thing2 is empty")
	}
	args = append(args, thing1)
	args = append(args, thing2)
	return name, args, nil
}
