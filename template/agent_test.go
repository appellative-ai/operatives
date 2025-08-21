package template

import (
	"fmt"
)

func ExampleNewAgent() {
	a := newAgent(nil)

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [common:core:agent/template/center]

}
