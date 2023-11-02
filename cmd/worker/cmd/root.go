package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Uonx/gather/internal/grpc/client"
	v1 "github.com/Uonx/gather/internal/proto/v1"
	"github.com/Uonx/gather/pkg/analyzer"
	"github.com/Uonx/gather/pkg/model"
	"github.com/spf13/cobra"
)

var (
	network           string
	address           string
	name              string
	scheduler_network string
	scheduler_address string
	rootCmd           = &cobra.Command{
		Use:   "scheduler",
		Short: "The scheduler of the collector",
		Long:  "Distribute tasks to jobs and report the results to the data receiver.",
		PreRun: func(cmd *cobra.Command, args []string) {
			// fmt.Println("scheduler prerun")
		},
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println(network, address, name, scheduler_network, scheduler_address)
			run()

			// conn, err := client.GrpcDial(network, address)
			// if err != nil {
			// 	panic(err)
			// }
			// defer conn.Close()
			// client := v1.NewTemplateTaskEventClient(conn)
			// task := v1.Task{
			// 	Url:      "https://www.instagram.com/api/v1/feed/user/xeesoxee/username/?count=12",
			// 	Auth:     `dpr=2; mid=ZTOuUwALAAFZemMLSNslIwly59Cg; ig_did=98D6B537-2358-4E5B-9472-7DAA9F88AB83; ig_nrcb=1; datr=UK4zZZMM3yqipnIN7TRbpT0Y; ds_user_id=62245146459; csrftoken=C2vlGHVfTyibWEoVuGtLWHa8rALEBfpT; sessionid=62245146459%3AGzclzHsIQYEBFJ%3A29%3AAYfQGdNe3q4PGdPGm2LqgIZA_k4d_vL7sul1v8r0lA; rur="NAO\05462245146459\0541729481311:01f750153cd8d7451713a915e71c94c5de2c0a4ee41f0f1c35dadd6704d2d0e7392ea31e"`,
			// 	Id:       "test1",
			// 	Template: "instagram_post",
			// 	Proxy:    "http://192.168.1.38:7890",
			// }
			// result, err := client.TemplateResult(context.Background(), &task)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// fmt.Println(result)
		},
	}
)

func run() {
	conn, err := client.GrpcDial(scheduler_network, scheduler_address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	scheduler_client := v1.NewSchedulerEventClient(conn)
	ticker := time.Tick(5 * time.Second)

	go func() {
		for {
			health := v1.Health{
				Name:   name,
				Status: 1,
			}
			_, err := scheduler_client.HealthResult(context.Background(), &health)
			if err != nil {
				fmt.Printf("conn scheduler %s error : %v \n", scheduler_address, err)
			}
			<-ticker
		}
	}()
	registry := v1.RegisterEvent{
		Work: &v1.Work{
			Name:      name,
			Ip:        "127.0.0.1", //获取本地ip地址
			Templates: []string{string(analyzer.TypeFaceBookPost), string(analyzer.TypeFaceBookUserPost), string(analyzer.TypeInstagramPost), string(analyzer.TypeInstagramPostComment)},
		},
	}

retry:
	stream, err := scheduler_client.Registry(context.Background(), &registry)
	if err != nil {
		time.Sleep(time.Second * 10)
		goto retry
	}
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			time.Sleep(time.Second)
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("recv data %s \n", recv.Data)
		task := model.Task{}
		err = json.Unmarshal([]byte(recv.Data), &task)
		if err != nil {
			fmt.Printf("unmarsharl %s error", recv.Data)
		}

		conn, err := client.GrpcDial(network, address)
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := v1.NewTemplateTaskEventClient(conn)

		t := v1.Task{
			Url:      task.ParentTask.Url,
			Auth:     task.Auth,
			Id:       task.Id,
			TaskId:   task.TaskId,
			Template: task.Template,
			Proxy:    task.Proxy,
			Params:   task.Params,
		}
		result, err := client.TemplateResult(context.Background(), &t)
		if err != nil {
			//做重试
			fmt.Println(err)
		}
		scheduler_client.WorkResult(context.Background(), result)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&network, "network", "unix", "unix")
	rootCmd.PersistentFlags().StringVar(&address, "address", "/tmp/template.sock", "/tmp/template.sock")
	rootCmd.PersistentFlags().StringVar(&name, "name", "work", "work")
	rootCmd.PersistentFlags().StringVar(&scheduler_network, "scheduler_network", "tcp", "tcp")
	rootCmd.PersistentFlags().StringVar(&scheduler_address, "scheduler_address", ":8000", ":8000")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
