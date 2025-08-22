package concurrency

import "fmt"

func ExampleNewAgent() {
	a := NewAgent(nil, 0)

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [common:core:agent/operative/concurrency]

}
