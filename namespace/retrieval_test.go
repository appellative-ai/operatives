package namespace

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/appellative-ai/center/template"
	"github.com/appellative-ai/postgres/retrieval"
	"net/http"
	"reflect"
	"testing"
)

func ExampleRetrieval() {
	r := tagRetrieval{
		Name: "test:agent",

		Args: []arg{
			{Name: "kind", Value: "aspect"},
			{Name: "count", Value: "25"},
		},
	}

	buf, err := json.Marshal(r)

	fmt.Printf("test: Retrieval() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Retrieval() -> [{"name":"test:agent","args":[{"name":"kind","value":"aspect"},{"name":"count","value":"25"}]}] [err:<nil>]

}

func Test_retrievalRequest(t *testing.T) {
	type args struct {
		ctx       context.Context
		retriever *retrieval.Interface
		processor template.Agent
		r         *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    bytes.Buffer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := retrievalRequest(tt.args.ctx, tt.args.retriever, tt.args.processor, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("retrievalRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("retrievalRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
