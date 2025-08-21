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

func ExampleRelation() {
	r := tagRelation{
		Name: "test:agent",
		//Instance: "core:aspect/resiliency",
		//Pattern:  "core:aspect/expressive",
		Args: []arg{
			{Name: "instance", Value: "core:aspect/resiliency"},
			{Name: "pattern", Value: "core:aspect/expressive"},
			{Name: "kind", Value: "aspect"},
			{Name: "count", Value: "25"},
		},
	}

	buf, err := json.Marshal(r)

	fmt.Printf("test: Relation() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Relation() -> [{"name":"test:agent","instance":"core:aspect/resiliency","pattern":"core:aspect/expressive","args":[{"name":"kind","value":"aspect"},{"name":"count","value":25}]}] [err:<nil>]

}

func Test_relationRequest(t *testing.T) {
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
			got, err := relationRequest(tt.args.ctx, tt.args.retriever, tt.args.processor, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("relationRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
