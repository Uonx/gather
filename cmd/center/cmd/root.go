package cmd

import (
	"fmt"
	"os"

	"github.com/Uonx/gather/pkg/db"
	"github.com/Uonx/gather/pkg/entities"
	"github.com/Uonx/gather/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
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
			mysqldb.AutoMigrate(&entities.Project{}, &entities.Task{}, &entities.SubTask{})
		},
		Run: func(cmd *cobra.Command, args []string) {
			g := gin.Default()
			g.POST("/project", Project)
			g.Run(address)
		},
	}
)

func Project(c *gin.Context) {
	project := model.Project{}
	c.BindJSON(&project)
	project_entities := entities.Project{
		Name:       project.Name,
		CronStatus: project.CronStatus,
		Cron:       project.Cron,
		Status:     model.Processing,
	}
	for _, v := range project.Task {
		task := entities.Task{
			Url:      v.Url,
			Auth:     v.Auth,
			Template: v.Template,
			Proxy:    v.Proxy,
			Status:   model.Processing,
		}
		task.SubTask = append(task.SubTask, entities.SubTask{
			Url:      v.Url,
			Auth:     v.Auth,
			Template: v.Template,
			Proxy:    v.Proxy,
			Status:   model.Unprocessed,
			Params:   "null",
		})
		project_entities.Task = append(project_entities.Task, task)
	}
	err := db.MysqlDB().Create(&project_entities).Error
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, project_entities.Id)
}

func init() {
	rootCmd.PersistentFlags().StringVar(&network, "network", "tcp", "tcp")
	rootCmd.PersistentFlags().StringVar(&address, "address", ":8008", "default :8008")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
