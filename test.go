package main

import (
	"context"
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
type UserAdd struct {
	UserId       int64 `mapstructure:"user_id"`
	RowsAffected int64 `mapstructure:"rows_affected"`
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

	paramBuilder := builder.NewParamBuilder().Add("name", "ruby").Add("age", 16)
	api := builder.NewApiBuilder("adduser", 1)

	userAdd := &UserAdd{}
	err := api.Invoke(context.Background(), c, paramBuilder, &userAdd)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(userAdd)
}
