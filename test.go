package main

import (
	"context"
	"fmt"
	"github.com/xuliangTang/istio-dbcore-sdk/pbfiles"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

type UserModel struct {
	Id   int64  `mapstructure:"id"`
	Name string `mapstructure:"name"`
	Age  uint8  `mapstructure:"age"`
}

func main() {
	client, _ := grpc.DialContext(context.Background(),
		"localhost:8080",
		grpc.WithInsecure(),
	)
	var c pbfiles.DBServiceClient
	c = pbfiles.NewDBServiceClient(client)

	/*structPb, _ := structpb.NewStruct(map[string]interface{}{
		"id": 9,
	})

	params := &pbfiles.SimpleParam{
		Params: structPb,
	}

	req := &pbfiles.QueryRequest{Name: "userlist", Params: params}
	rsp, _ := c.Query(context.Background(), req)
	for _, item := range rsp.Result {
		fmt.Println(item.AsMap())
	}*/

	paramBuilder := builder.NewParamBuilder().Add("id", 9)
	api := builder.NewApiBuilder("userlist", 0)

	users := make([]*UserModel, 0)
	err := api.Invoke(context.Background(), c, paramBuilder, &users)
	if err != nil {
		log.Fatalln(err)
	}

	for _, u := range users {
		fmt.Println(u)
	}
}
