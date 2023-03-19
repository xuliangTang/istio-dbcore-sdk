package main

import (
	"context"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
	"time"
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

func main() {
	// 客户端构建器
	c, err := builder.NewClientBuilder("localhost:8080").WithOptions(grpc.WithInsecure()).Build()
	if err != nil {
		log.Fatalln(err)
	}

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

	log.Println(userAdd)
}
