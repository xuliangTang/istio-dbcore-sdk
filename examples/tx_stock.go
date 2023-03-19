package examples

import (
	"context"
	"fmt"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

type ProdStock struct {
	ProdID  int `mapstructure:"prod_id"`
	Stock   int `mapstructure:"stock"`
	Version int `mapstructure:"version"`
}

func TestTxStock(pid, num int) {
	client, err := builder.NewClientBuilder("localhost:8080").WithOptions(grpc.WithInsecure()).Build()
	if err != nil {
		log.Fatal(err)
	}

	// 创建事务API
	txApi := builder.NewTxApi(context.Background(), client)
	err = txApi.Tx(func(tx *builder.TxApi) error {
		ps := &ProdStock{}
		psParam := builder.NewParamBuilder().Add("prodId", pid)
		err := tx.QueryForModel("getstock", psParam, ps)
		if err != nil {
			return err
		}
		log.Println("getstock成功", ps.ProdID, ps.Stock)

		// time.Sleep(time.Second * 2)

		if ps.Stock < num {
			return fmt.Errorf("库存不够了")
		}

		setStockParam := builder.NewParamBuilder().Add("prodId", pid).Add("stock", ps.Stock-num).Add("version", ps.Version)
		execRet := &ExecResult{} // 增删改执行结果
		err = tx.Exec("setstock", setStockParam, execRet)
		if err != nil || execRet.RowsAffected == 0 {
			return fmt.Errorf("扣减库存失败")
		}
		log.Println("setstock成功", execRet.RowsAffected)

		return nil
	})

	log.Println(err)
}
