package cmd

import (
	"fmt"
	"os"

	"github.com/Uonx/gather/internal/grpc/server"
	v1 "github.com/Uonx/gather/internal/proto/v1"
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
			// fmt.Println("scheduler prerun")
		},
		Run: func(cmd *cobra.Command, args []string) {
			server.StartGrpcServer(network, address, registryService)
		},
	}
)

func registryService(s *grpc.Server) func() {
	return func() {
		v1.RegisterTemplateTaskEventServer(s, &service.TemplateTaskService{})
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&network, "network", "unix", "unix")
	rootCmd.PersistentFlags().StringVar(&address, "address", "/tmp/template.sock", "/tmp/template.sock")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
