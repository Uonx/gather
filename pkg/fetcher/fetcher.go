package fetcher

import (
	"context"

	"github.com/Uonx/gather/pkg/analyzer"
	_ "github.com/Uonx/gather/pkg/analyzer/all"
	"github.com/Uonx/gather/pkg/model"
)

func Analyzer(ctx context.Context, task model.Task) (model.AnalyzerResult, error) {
	analyzer := analyzer.Analyzers[analyzer.Type(task.Template)]
	req := model.Request{
		Url:    task.Url,
		Method: task.Method,
		Proxy:  task.Proxy,
		Params: task.Params,
		ParentRequest: model.ParentRequest{
			Url:    task.ParentTask.Url,
			Method: task.ParentTask.Method,
			Proxy:  task.ParentTask.Proxy,
		},
	}
	headers, err := analyzer.SetHeader(task.Auth, &req)
	if err != nil {
		return model.AnalyzerResult{}, err
	}
	req.Headers = headers
	content, err := analyzer.Fetcher(&req)

	if err != nil {
		return model.AnalyzerResult{}, err
	}
	return analyzer.Analyze(ctx, content, task, req)
}
