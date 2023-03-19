package examples

import (
	"context"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
	"time"
)

type userModel struct {
	Id   int64  `mapstructure:"id"`
	Name string `mapstructure:"name"`
	Age  uint8  `mapstructure:"age"`
}

func TestQuery() {
	// 客户端构建器
	c, err := builder.NewClientBuilder("localhost:8080").WithOptions(grpc.WithInsecure()).Build()
	if err != nil {
		log.Fatalln(err)
	}

	// 参数构建器
	paramBuilder := builder.NewParamBuilder().Add("id", 50)

	// api构建器
	api := builder.NewApiBuilder("userlist", builder.ApiTypeQuery)

	// 查询结果集
	users := make([]*userModel, 0)

	// 执行调用API
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()
	err = api.Invoke(ctx, c, paramBuilder, &users)
	if err != nil {
		log.Fatalln(err)
	}

	for _, u := range users {
		log.Println(u)
	}
}
