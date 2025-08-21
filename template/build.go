package template

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Build(args []Arg, params []Param) ([]any, error) {
	if len(args) == 0 {
		return nil, errors.New("input args are empty")
	}
	if len(params) == 0 {
		return nil, errors.New("input parameters are empty")
	}
	slices.SortFunc(params, func(a, b Param) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	slices.SortFunc(args, func(a, b Arg) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	var ix = 0
	var temp []any

	for _, a := range args {
		p := params[ix]
		switch strings.Compare(a.Name, p.Name) {
		case 0:
			v, err := createValue(a, p)
			if err != nil {
				return nil, err
			}
			temp = append(temp, v)
			ix++
		case -1:
			return nil, errors.New(fmt.Sprintf("argument name [%v] is not in the parameter list", a.Name))
		case 1:
			if !p.Nullable {
				return nil, errors.New(fmt.Sprintf("parameter [%v] does not allow nulls", p.Name))
			}
			ix++
		}
		if ix >= len(params) {
			return nil, errors.New(fmt.Sprintf("argument name [%v] is not in the parameter list", a.Name))
		}
	}
	return temp, nil
}

func createValue(a Arg, p Param) (any, error) {
	switch p.Type {
	case "string":
		return a.Value, nil
	case "int":
		i, err := strconv.Atoi(a.Value)
		return i, err
	}
	return nil, errors.New(fmt.Sprintf("parameter [%v] type [%v] is not supported", p.Name, p.Type))
}
