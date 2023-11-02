package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/Uonx/gather/internal/grpc/client"
	v1 "github.com/Uonx/gather/internal/proto/v1"
	"github.com/Uonx/gather/pkg/model"
)

func TestTemplate(t *testing.T) {
	network := "unix"
	address := "/tmp/template.sock"

	conn, err := client.GrpcDial(network, address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	template := v1.NewTemplateTaskEventClient(conn)
	task := model.Task{
		Id:       116,
		TaskId:   53,
		Proxy:    "http://192.168.1.26:7890",
		Url:      "https://www.facebook.com/bbc",
		Template: "facebook_user_post",
	}
	t1 := v1.Task{
		Url:      task.Url,
		Auth:     task.Auth,
		Id:       task.Id,
		TaskId:   task.TaskId,
		Template: task.Template,
		Proxy:    task.Proxy,
		Params:   task.Params,
	}
	result, err := template.TemplateResult(context.Background(), &t1)
	if err != nil {
		//做重试
		fmt.Println(err)
	}

	fmt.Println(result)
}
