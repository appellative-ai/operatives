package templatetest

import "fmt"

const (
	fileName = "file://[cwd]/template.json"
	name3    = "common:core:retrieval/query/test3"
	name2    = "common:core:retrieval/get/test2"
)

func ExampleNewAgent() {
	agent := newAgent(fileName)
	fmt.Printf("test: newAgent() -> [%v]\n", agent)

	t, ok := agent.cache.Load(name3)
	fmt.Printf("test: Entry() -> [%v] [ok:%v]\n", t, ok)

	t, ok = agent.cache.Load(name2)
	fmt.Printf("test: Entry() -> [%v] [ok:%v]\n", t, ok)

	//Output:
	//test: newAgent() -> [common:core:agent/template/center/test]
	//test: Entry() -> [{common:core:retrieval/query/test3 CALL dbo.QueryNamespace($1,$2,$3) [{name true string } {count false int } {createDate false string DateTime}]}] [ok:true]
	//test: Entry() -> [{common:core:retrieval/get/test2 CALL dbo.GetThing($1,$2) [{name true string } {count false int }]}] [ok:true]

}
