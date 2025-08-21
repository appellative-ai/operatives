package namespace

import (
	"fmt"
	"github.com/appellative-ai/core/std"
	"github.com/appellative-ai/postgres/request/requesttest"
	"net/url"
)

func ExampleCreateLinkArgs() {
	values := make(url.Values)

	name, args, err := createLinkArgs(values)
	fmt.Printf("test: createLinkArgs() -> [name:%v] [args:%v] [err:%v]\n", name, args, err)

	values.Add(nameName, "common:resiliency:agent/rate-limiting/request/http")
	name, args, err = createLinkArgs(values)
	fmt.Printf("test: createLinkArgs() -> [name:%v] [args:%v] [err:%v]\n", name, args, err)

	values.Add(authorName, "bobs uncle")
	name, args, err = createLinkArgs(values)
	fmt.Printf("test: createLinkArgs() -> [name:%v] [args:%v] [err:%v]\n", name, args, err)

	values.Add(thing1Name, "common:core:agent/thing1")
	name, args, err = createLinkArgs(values)
	fmt.Printf("test: createLinkArgs() -> [name:%v] [args:%v] [err:%v]\n", name, args, err)

	values.Add(thing2Name, "common:core:agent/thing2")
	name, args, err = createLinkArgs(values)
	fmt.Printf("test: createLinkArgs() -> [name:%v] [args:%v] [err:%v]\n", name, args, err)

	//Output:
	//test: createLinkArgs() -> [name:] [args:[]] [err:name is empty]
	//test: createLinkArgs() -> [name:] [args:[]] [err:author is empty]
	//test: createLinkArgs() -> [name:] [args:[]] [err:thing1 is empty]
	//test: createLinkArgs() -> [name:] [args:[]] [err:thing2 is empty]
	//test: createLinkArgs() -> [name:common:resiliency:agent/rate-limiting/request/http] [args:[common:resiliency:agent/rate-limiting/request/http bobs uncle common resiliency agent /rate-limiting/request/http common:core:agent/thing1 common:core:agent/thing2]] [err:<nil>]

}

func ExampleLinkRequest() {
	name := "common:resiliency:agent/rate-limiting/request/http"

	r, err := linkRequest(nil, nil, nil)
	fmt.Printf("test: linkRequest() -> [result:%v] [err:%v]\n", r, err)

	m := std.NewSyncMap[string, any]()
	requester := requesttest.NewRequester(m)
	values := make(url.Values)
	values.Add(nameName, name)
	values.Add(authorName, "bobs uncle")
	values.Add(thing1Name, "common:core:agent/thing1")
	values.Add(thing2Name, "common:core:agent/thing2")

	r, err = linkRequest(nil, requester, values)
	fmt.Printf("test: linkRequest() -> [result:%v] [err:%v]\n", r, err)

	t, ok := m.Load(name)
	fmt.Printf("test: Load() -> [%v] [ok:%v]\n", t, ok)

	//Output:
	//test: linkRequest() -> [result:{ 0 false false false false}] [err:query values are nil]
	//test: linkRequest() -> [result:{ 1 true false false false}] [err:<nil>]
	//test: Load() -> [common:resiliency:agent/rate-limiting/request/http] [ok:true]

}
