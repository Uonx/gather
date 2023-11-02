package service

import (
	"context"

	"github.com/Uonx/gather/internal/proto/basic"
	v1 "github.com/Uonx/gather/internal/proto/v1"
	"github.com/Uonx/gather/pkg/fetcher"
	"github.com/Uonx/gather/pkg/model"
)

type TemplateTaskService struct {
	v1.UnimplementedTemplateTaskEventServer
}

func (TemplateTaskService) TemplateResult(ctx context.Context, t *v1.Task) (*v1.Result, error) {
	task := model.Task{
		Id:       t.Id,
		Url:      t.Url,
		Auth:     t.Auth,
		Proxy:    t.Proxy,
		TaskId:   t.TaskId,
		Template: t.Template,
		Params:   t.Params,
	}
	resutl, err := fetcher.Analyzer(context.Background(), task)
	r := v1.Result{
		Response: &basic.Response{
			Error: &basic.Error{},
		},
		CurrentTask: &v1.Task{
			Id:       task.Id,
			TaskId:   task.TaskId,
			Url:      task.Url,
			Auth:     task.Auth,
			Proxy:    task.Proxy,
			Template: task.Template,
		},
	}
	if err != nil {
		r.Response.Error = &basic.Error{Code: 400, Message: err.Error()}
		return &r, nil
	}
	r.Response.Error = &basic.Error{Code: 200, Message: "ok"}
	for _, v := range resutl.Items {
		r.Items = append(r.Items, &v1.Item{
			Url:     v.Url,
			PayLoad: v.PayLoad,
			Task: &v1.Task{
				TaskId:   v.Task.TaskId,
				Url:      v.Task.Url,
				Auth:     v.Task.Auth,
				Proxy:    v.Task.Proxy,
				Template: v.Task.Template,
				Params:   v.Task.Params,
			},
		})
	}
	for _, v := range resutl.Tasks {
		r.Tasks = append(r.Tasks, &v1.Task{
			Url:      v.Url,
			Id:       v.TaskId,
			Auth:     v.Auth,
			Proxy:    v.Auth,
			Template: v.Template,
			Params:   v.Params,
		})
	}
	return &r, nil
}
