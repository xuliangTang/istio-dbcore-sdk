package builder

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/xuliangTang/istio-dbcore-sdk/pbfiles"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/helpers"
)

const (
	ApiTypeQuery = iota
	ApiTypeExec
)

type ApiBuilder struct {
	name    string // api名称
	apiType uint8
}

func NewApiBuilder(name string, apiType uint8) *ApiBuilder {
	return &ApiBuilder{name: name, apiType: apiType}
}

// Invoke 普通执行
func (this *ApiBuilder) Invoke(ctx context.Context, client pbfiles.DBServiceClient, builder *ParamBuilder, out interface{}) error {
	if this.apiType == ApiTypeQuery {
		req := &pbfiles.QueryRequest{Name: this.name, Params: builder.Build()}
		rsp, err := client.Query(ctx, req)
		if err != nil {
			return err
		}

		if out != nil {
			mapList := helpers.PbStructToMapList(rsp.Result)
			return mapstructure.Decode(mapList, out)
		}

		return nil

	} else if this.apiType == ApiTypeExec {
		req := &pbfiles.ExecRequest{Name: this.name, Params: builder.Build()}
		rsp, err := client.Exec(ctx, req)
		if err != nil {
			return err
		}

		var m map[string]interface{}
		if out != nil {
			if rsp.Select != nil {
				m = rsp.Select.AsMap()
				m["rows_affected"] = rsp.RowsAffected
			}
		} else {
			m = map[string]interface{}{"rows_affected": rsp.RowsAffected}
		}

		return mapstructure.Decode(m, out)
	}

	return nil
}
