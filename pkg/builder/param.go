package builder

import (
	"github.com/xuliangTang/istio-dbcore-sdk/pbfiles"
	"google.golang.org/protobuf/types/known/structpb"
)

type ParamBuilder struct {
	params map[string]interface{}
}

func NewParamBuilder() *ParamBuilder {
	return &ParamBuilder{params: make(map[string]interface{})}
}

func (this *ParamBuilder) Add(name string, value interface{}) *ParamBuilder {
	this.params[name] = value
	return this
}

func (this *ParamBuilder) Build() *pbfiles.SimpleParam {
	paramStruct, _ := structpb.NewStruct(this.params)
	return &pbfiles.SimpleParam{
		Params: paramStruct,
	}
}
