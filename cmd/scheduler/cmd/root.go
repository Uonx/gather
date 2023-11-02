package cmd

import (
	"fmt"
	"os"

	"github.com/Uonx/gather/internal/grpc/server"
	v1 "github.com/Uonx/gather/internal/proto/v1"
	"github.com/Uonx/gather/pkg/db"
	"github.com/Uonx/gather/pkg/entities"
	"github.com/Uonx/gather/pkg/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	network string
	address string
	rootCmd = &cobra.Command{
		Use:   "scheduler",
		Short: "The scheduler of the collector",
		Long:  "Distribute tasks to jobs and report the results to the data receiver.",
		PreRun: func(cmd *cobra.Command, args []string) {
			mysql := db.MysqlOpts{
				Endpoint: "127.0.0.1:3306",
				Username: "root",
				Password: "test123456",
				Database: "gather",
			}
			mysqldb, err := db.NewMysqlClient(&mysql)
			if err != nil {
				panic(err)
			}
			mysqldb.AutoMigrate(&entities.Project{}, &entities.Task{}, &entities.SubTask{}, &entities.Result{})
		},
		Run: func(cmd *cobra.Command, args []string) {
			s := service.SchedulerService{}
			go s.SendMessage()
			// go service.SendMessage()
			server.StartGrpcServer(network, address, registryService)
		},
	}
)

func registryService(s *grpc.Server) func() {
	return func() {
		v1.RegisterSchedulerEventServer(s, &service.SchedulerService{})
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&network, "network", "tcp", "tcp")
	rootCmd.PersistentFlags().StringVar(&address, "address", ":8000", "default :8000")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
