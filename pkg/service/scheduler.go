package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/Uonx/gather/internal/proto/basic"
	v1 "github.com/Uonx/gather/internal/proto/v1"
	"github.com/Uonx/gather/pkg/db"
	"github.com/Uonx/gather/pkg/entities"
	"github.com/Uonx/gather/pkg/model"
	"github.com/Uonx/gather/pkg/utils"
)

type SchedulerService struct {
	v1.UnimplementedSchedulerEventServer
}

func (s *SchedulerService) Registry(data *v1.RegisterEvent, stream v1.SchedulerEvent_RegistryServer) error {
	if data.GetWork() == nil ||
		len(data.GetWork().GetName()) == 0 ||
		len(data.GetWork().Templates) == 0 ||
		len(data.GetWork().Ip) == 0 {
		return fmt.Errorf("work not set")
	}

	wd := utils.Data{
		Name:      data.Work.Name,
		Ip:        data.Work.Ip,
		Templates: data.Work.Templates,
		Stream:    &stream,
	}
	utils.WorkDB.AddData(&wd)
	streamCtx := (*wd.Stream).Context()
	for {
		select {
		case <-streamCtx.Done():
			utils.WorkDB.DeleteData(wd)
			return nil
		}
	}

}

var ticker = time.Tick(3 * time.Second)

func (s *SchedulerService) SendMessage() {
	utils.NewWorkDb(0) //初始化work db
	for {
		<-ticker
		var subTasks []entities.SubTask
		err := db.MysqlDB().Model(entities.SubTask{}).Preload("Task").Where(entities.SubTask{Status: model.Unprocessed}).Limit(10).Find(&subTasks).Error
		if err != nil {
			fmt.Printf("find subtask error: %v \n", err)
		}
		wg := sync.WaitGroup{}
		for _, v := range subTasks {
			wg.Add(1)
			go func(v entities.SubTask) {
				task := model.Task{
					Id:       int64(v.Id),
					Url:      v.Url,
					Auth:     v.Auth,
					TaskId:   int64(v.TaskId),
					Template: v.Template,
					Proxy:    v.Proxy,
					Params:   v.Params,
					ParentTask: model.ParentTask{
						Id:       int64(v.Task.Id),
						Url:      v.Task.Url,
						Auth:     v.Task.Auth,
						TaskId:   int64(v.Task.ProjectId),
						Template: v.Task.Template,
					},
				}

				taskbyte, _ := json.Marshal(&task)
				var data utils.Data
				var ok bool = false
				for !ok {
					data, ok = utils.WorkDB.Borrow(v.Template)
				}
				stream := data.Stream
				err := (*stream).Send(&v1.Execution{
					Data: string(taskbyte),
				})
				// time.Sleep(5 * time.Second)
				utils.WorkDB.Return(data)
				if err != nil {
					fmt.Printf("task send error: %v \n", err)
					wg.Done()
					return
				}
				v.Status = model.Processing
				err = db.MysqlDB().Save(&v).Error
				if err != nil {
					fmt.Printf("task send error: %v \n", err)
					wg.Done()
					return
				}
				wg.Done()
			}(v)
		}
		wg.Wait()
	}
}

func (s *SchedulerService) HealthResult(health *v1.Health, stream v1.SchedulerEvent_HealthResultServer) error {
	// fmt.Println(health)
	// 做work删除
	return nil
}

func (s *SchedulerService) WorkResult(ctx context.Context, result *v1.Result) (*basic.Response, error) {
	// result.Response    //error code ==200成功
	// result.CurrentTask // 当前任务
	// result.Items       //返回的每条结果
	// result.Tasks       //需要继续的任务
	// fmt.Println(result)
	subTask := entities.SubTask{}
	err := db.MysqlDB().Model(entities.SubTask{}).Find(&subTask, result.CurrentTask.Id).Error
	if err != nil {
		fmt.Printf("find subtask %s error: %v\n", result.CurrentTask.Id, err)
	}

	if result.Response.Error.Code == 200 {
		subTask.Status = model.Completed
	} else {
		subTask.Status = model.Failed
		subTask.ErrMessage = result.Response.Error.Message
	}
	err = db.MysqlDB().Save(&subTask).Error
	if err != nil {
		fmt.Printf("change subtask %s status error: %v\n", result.CurrentTask.Id, err)
	}
	subTasks := []entities.SubTask{}
	for _, v := range result.Tasks {
		subTasks = append(subTasks, entities.SubTask{
			TaskId:   int(result.CurrentTask.TaskId),
			Url:      v.Url,
			Auth:     v.Auth,
			Template: v.Template,
			Proxy:    result.CurrentTask.Proxy,
			Status:   model.Unprocessed,
			Params:   v.Params,
		})
	}
	if len(subTasks) > 0 {
		err = db.MysqlDB().Create(&subTasks).Error
		if err != nil {
			fmt.Printf("save result.tasks error: %v\n", err)
		}
	}
	// results := []entities.Result{}
	for _, v := range result.Items {
		// results = append(results, entities.Result{
		// 	TaskId:  v.Task.TaskId,
		// 	SubId:   v.Task.Id,
		// 	PayLoad: v.PayLoad,
		// })
		fmt.Println(v)
	}
	// if len(results) > 0 {
	// 	err = db.MysqlDB().Create(&results).Error
	// 	if err != nil {
	// 		fmt.Printf("save result.results error: %v\n", err)
	// 	}
	// }
	return &basic.Response{
		Error: &basic.Error{
			Code: 200,
		},
	}, nil
}
