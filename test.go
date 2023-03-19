package main

import (
	"context"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

type UserModel struct {
	Id   int64  `mapstructure:"id"`
	Name string `mapstructure:"name"`
	Age  uint8  `mapstructure:"age"`
}
type UserAdd struct {
	UserId       int64 `mapstructure:"user_id"`
	RowsAffected int64 `mapstructure:"rows_affected"`
}
type UserScoreAdd struct {
	RowsAffected int64 `mapstructure:"rows_affected"`
}

func main() {
	// 客户端构建器
	c, err := builder.NewClientBuilder("localhost:8080").WithOptions(grpc.WithInsecure()).Build()
	if err != nil {
		log.Fatalln(err)
	}

	txApi := builder.NewTxApi(context.Background(), c)
	err = txApi.Tx(func(tx *builder.TxApi) error {
		paramBuilder := builder.NewParamBuilder().Add("name", "hua").Add("age", 23)
		user := &UserAdd{}
		err := tx.Exec("adduser", paramBuilder, user)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("addUser成功", user.UserId, user.RowsAffected)

		paramBuilder = builder.NewParamBuilder().Add("userId", user.UserId).Add("score", 66)
		userScore := &UserScoreAdd{}
		err = tx.Exec("adduserscore", paramBuilder, userScore)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("addUserScore成功", userScore.RowsAffected)
		return nil
	})
	log.Println(err)

	/*
		// 参数构建器
		paramBuilder := builder.NewParamBuilder().Add("name", "ruby").Add("age", 16)

		// api构建器
		api := builder.NewApiBuilder("adduser", 1)

		// 调用
		userAdd := &UserAdd{}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		err = api.Invoke(ctx, c, paramBuilder, &userAdd)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(userAdd)*/
}
