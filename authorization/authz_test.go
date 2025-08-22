package authorization

import (
	"fmt"
	"github.com/appellative-ai/common/core"
)

func ExampleAuthorization_Chain() {
	//name := "agent/authorization"
	chain := core.BuildNetwork([]any{Authorization})
	fmt.Printf("test: BuildExchangeChain() -> %v\n", chain != nil)

	//exchange.RegisterExchangeHandler(name, Authorization)
	//l := exchange.ExchangeHandler(name)
	//fmt.Printf("test: ExchangeLink() -> %v %v\n", reflect.TypeOf(Authorization), reflect.TypeOf(l))
	//chain = rest.BuildNetwork([]any{l})
	//fmt.Printf("test: repository.ExchangeLink() -> %v\n", chain != nil)

	//Output:
	//test: BuildExchangeChain() -> true
	
}
