package examples

import (
	"context"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

type addUserScore struct {
	RowsAffected int64 `mapstructure:"rows_affected"`
}

func TestTx() {
	// 客户端构建器
	c, err := builder.NewClientBuilder("localhost:8080").WithOptions(grpc.WithInsecure()).Build()
	if err != nil {
		log.Fatalln(err)
	}

	// 创建事务API
	txApi := builder.NewTxApi(context.Background(), c)
	err = txApi.Tx(func(tx *builder.TxApi) error {
		// 创建用户
		paramBuilder := builder.NewParamBuilder().Add("name", "hua").Add("age", 23)
		addUserRet := &addUser{}
		err := tx.Exec("adduser", paramBuilder, addUserRet)
		if err != nil {
			return err
		}
		log.Println("addUser成功", addUserRet.UserId, addUserRet.RowsAffected)

		// 给用户赠送积分
		paramBuilder = builder.NewParamBuilder().Add("userId", addUserRet.UserId).Add("score", 66)
		addUserScoreRet := &addUserScore{}
		err = tx.Exec("adduserscore", paramBuilder, addUserScoreRet)
		if err != nil {
			return err
		}
		log.Println("addUserScore成功", addUserScoreRet.RowsAffected)
		return nil
	})

	log.Println(err)
}
