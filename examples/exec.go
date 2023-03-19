package examples

import (
	"context"
	"fmt"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

type addUser struct {
	UserId       int64 `mapstructure:"user_id"`
	RowsAffected int64 `mapstructure:"rows_affected"`
}

func TestExec() {
	// 客户端构建器
	c, err := builder.NewClientBuilder("localhost:8080").WithOptions(grpc.WithInsecure()).Build()
	if err != nil {
		log.Fatalln(err)
	}

	// 参数构建器
	paramBuilder := builder.NewParamBuilder().Add("name", "zs").Add("age", 18)

	// api构建器
	api := builder.NewApiBuilder("adduser", builder.ApiTypeExec)

	// 查询结果集
	addUserRet := &addUser{}

	// 执行调用API
	err = api.Invoke(context.Background(), c, paramBuilder, addUserRet)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addUserRet)
}
