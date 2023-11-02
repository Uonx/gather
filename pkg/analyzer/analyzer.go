package analyzer

import (
	"context"
	"fmt"

	"github.com/Uonx/gather/pkg/model"
)

var (
	Analyzers = make(map[Type]analyzer)
)

type analyzer interface {
	Type() Type
	Analyze(ctx context.Context, input []byte, task model.Task, req model.Request) (model.AnalyzerResult, error)
	Fetcher(req *model.Request) ([]byte, error)
	SetHeader(auth string, req *model.Request) (map[string]string, error)
}

func RegistryAnalyzer(analyzer analyzer) {
	if _, ok := Analyzers[analyzer.Type()]; ok {
		fmt.Printf("analyzer %s is registered twice", analyzer.Type())
	}
	Analyzers[analyzer.Type()] = analyzer
}
