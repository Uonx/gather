package fetcher

import (
	"context"
	"fmt"
	"testing"

	"github.com/Uonx/gather/pkg/analyzer"
	"github.com/Uonx/gather/pkg/model"
)

func TestAnalyzer(t *testing.T) {

	var cookie = `dpr=2; mid=ZTeEFwALAAFRWnThxDiVeTft_IAX; ig_did=5986B940-1B01-4C7C-8608-93E8A9ABBF29; ig_nrcb=1; datr=D4Q3ZZ12I3ZvFUV2qm-okpz2; csrftoken=2oMAvrjN0w8vned6TnOXzymzcsZ7pTKD; ds_user_id=62245146459; sessionid=62245146459%3Av5WVrsBxEA6rSN%3A20%3AAYeZVTRiFpZgMCm6DserbxPO6tzpoAeXr_hkseAkmg; rur="ODN\05462245146459\0541729673214:01f777df88ae0ca99a2583c7253b213cb5725e28a9f1dd1c0598fdeaac882c6e904e5328"`

	task := model.Task{
		Url:   "https://www.instagram.com/api/v1/feed/user/xeesoxee/username/?count=12",
		Auth:  cookie,
		Proxy: "http://192.168.1.38:7890",
		// TaskId:   "test1",
		Template: string(analyzer.TypeInstagramPost),
	}
	resutl, err := Analyzer(context.Background(), task)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resutl)
}
