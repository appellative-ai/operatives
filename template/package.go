package template

type Arg struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Result struct {
	Sql  string
	Args []any
}

type Param struct {
	Name     string `json:"name"`
	Nullable bool   `json:"nullable"`
	Type     string `json:"type"`
	SqlType  string `json:"sql-type"`
}

type Entry struct {
	Name   string  `json:"name"`
	Sql    string  `json:"sql"`
	Params []Param `json:"params"`
}
