package template

import (
	"encoding/json"
	"fmt"
)

func ExampleParams() {
	a := Entry{Name: "CALL dbo.QueryNamespace($1,$2,$3)", Params: []Param{
		{Name: "name", Nullable: true, Type: "string", SqlType: ""},
		{Name: "count", Nullable: false, Type: "int", SqlType: ""},
		{Name: "createDate", Nullable: false, Type: "string", SqlType: "DateTime"},
	},
	}

	fmt.Printf("test: template() -> [sql:%v] [args:%v]\n", a.Name, a.Params)
	buf, err := json.Marshal(a)
	fmt.Printf("test: template() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: template() -> [sql:CALL dbo.QueryNamespace($1,$2,$3)] [args:[{name true string } {count false int } {createDate false string DateTime}]]
	//test: template() -> [{"name":"CALL dbo.QueryNamespace($1,$2,$3)","args":[{"name":"name","nullable":true,"type":"string","sql-type":""},{"name":"count","nullable":false,"type":"int","sql-type":""},{"name":"createDate","nullable":false,"type":"string","sql-type":"DateTime"}]}] [err:<nil>]

}

/*
func ExampleArgs2() {
	a := template{Sql: "sp.QueryNamespace", Args: []arg{
		{Name: "name", Value: "test:agent"},
		{Name: "count", Value: 123},
	},
	}
	fmt.Printf("test: template() -> [sql:%v] [args:%v]\n", a.Sql, a.Args)
	buf, err := json.Marshal(a)
	fmt.Printf("test: template() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//fail
}


*/
