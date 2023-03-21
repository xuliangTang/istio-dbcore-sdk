package builder

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/xuliangTang/istio-dbcore-sdk/pbfiles"
	"google.golang.org/grpc"
)

type TxApi struct {
	ctx    context.Context
	cancel context.CancelFunc
	client pbfiles.DBService_TxClient
}

func (this *TxApi) GetClient() pbfiles.DBService_TxClient {
	return this.client
}

func NewTxApi(ctx context.Context, client pbfiles.DBServiceClient, opts ...grpc.CallOption) *TxApi {
	apiCtx, cancel := context.WithCancel(ctx)
	txClient, err := client.Tx(apiCtx, opts...)
	if err != nil {
		panic(err)
	}
	return &TxApi{ctx: ctx, client: txClient, cancel: cancel}
}

func (this *TxApi) Exec(apiName string, paramBuilder *ParamBuilder, out interface{}) error {
	err := this.client.Send(&pbfiles.TxRequest{Name: apiName, Params: paramBuilder.Build(), Type: "exec"})
	if err != nil {
		return err
	}

	rsp, err := this.client.Recv()
	if err != nil {
		return err
	}

	if out != nil {
		if execRet, ok := rsp.Result.AsMap()["exec"]; ok { // 返回[]interface{} 0:受影响的行 1:selectKey
			if execRet.([]interface{})[1] != nil {
				m := execRet.([]interface{})[1].(map[string]interface{})
				m["rows_affected"] = execRet.([]interface{})[0]
				return mapstructure.WeakDecode(m, out)
			} else {
				m := map[string]interface{}{"rows_affected": execRet.([]interface{})[0]}
				return mapstructure.WeakDecode(m, out)
			}
		}
	}

	return nil
}

func (this *TxApi) Query(apiName string, paramBuilder *ParamBuilder, out interface{}) error {
	err := this.client.Send(&pbfiles.TxRequest{Name: apiName, Params: paramBuilder.Build(), Type: "query"})
	if err != nil {
		return err
	}

	rsp, err := this.client.Recv()
	if err != nil {
		return err
	}

	if out != nil {
		if queryRet, ok := rsp.Result.AsMap()["query"]; ok { // 返回map[key]value
			return mapstructure.WeakDecode(queryRet, out)
		}
	}

	return nil
}

func (this *TxApi) QueryForModel(apiName string, paramBuilder *ParamBuilder, out interface{}) error {
	err := this.client.Send(&pbfiles.TxRequest{Name: apiName, Params: paramBuilder.Build(), Type: "query"})
	if err != nil {
		return err
	}

	rsp, err := this.client.Recv()
	if err != nil {
		return err
	}

	if out != nil {
		if queryRet, ok := rsp.Result.AsMap()["query"]; ok { // 返回map[key]value
			if retForMap, ok := queryRet.([]interface{}); ok && len(retForMap) == 1 {
				return mapstructure.WeakDecode(retForMap[0], out)
			} else {
				return fmt.Errorf("error query model: no result ")
			}
		}
	}

	return nil
}

func (this *TxApi) Tx(fn func(tx *TxApi) error) error {
	err := fn(this)
	if err != nil {
		this.cancel()
		return err
	}

	return this.client.CloseSend()
}
