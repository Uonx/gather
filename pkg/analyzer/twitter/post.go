package twitter

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"

	"github.com/Uonx/gather/pkg/analyzer"
	"github.com/Uonx/gather/pkg/model"
	"github.com/Uonx/gather/pkg/utils"
)

func init() {
	analyzer.RegistryAnalyzer(&twitterPostAnalyzer{})
}

type twitterPostAnalyzer struct{}

var headerParam = regexp.MustCompile(`ct0=([0-9A-Za-z]+)`)

func (i twitterPostAnalyzer) SetHeader(auth string, req *model.Request) (map[string]string, error) {
	headers := make(map[string]string)
	// headers["authority"] = "twitter.com"
	headers["Accept"] = "*/*"
	headers["Accept-Encoding"] = "gzip, deflate, br"
	headers["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6"
	headers["Authorization"] = `Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA`
	headers["Cache-Control"] = "no-cache"
	headers["Content-Type"] = "application/json"
	headers["Pragma"] = "no-cache"
	headers["Referer"] = "https://twitter.com/BillGates"
	headers["Sec-Ch-Ua"] = `"Chromium";v="118", "Microsoft Edge";v="118", "Not=A?Brand";v="99"`
	headers["Sec-Ch-Ua-Mobile"] = "?0"
	headers["Sec-Ch-Ua-Platform"] = `"Windows"`
	headers["Sec-Fetch-Dest"] = "empty"
	headers["Sec-Fetch-Mode"] = "cors"
	headers["Sec-Fetch-Site"] = "same-origin"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.61"
	headers["Cookie"] = auth
	// headers["X-Client-Transaction-Id"] = "qGojWVt6DrWbyYrxGyTunQYhvqzp//wNzlDrN64mrz+eEJ5pBt2VXIVJ/9zXaDhSvU5qTqjYm7+5LEVig9xmUQgQH1FHqQ"
	headers["X-Client-Uuid"] = "dc5b6b4f-7915-47a7-9feb-73aba3a8e33a"
	// params := headerParam.FindSubmatch([]byte(cookie))
	headers["X-Csrf-Token"] = "51dc3c138f0fdee86d331cb39afe97192e6cbaaa10f3350cbb0834fcc5e2a744a4ad381d4bd3b52dd6645c93682a0167934abffd8bda506f17d03519520decb4d1fea16a2ac34f2e1883714372016870" //string(params[1])
	headers["X-Twitter-Active-User"] = "yes"
	headers["X-Twitter-Auth-Type"] = "OAuth2Session"
	headers["X-Twitter-Client-Language"] = "en"
	headers["X-Tfe-User-Id"] = "1713138801609203712"
	return headers, nil
}

func (i twitterPostAnalyzer) Fetcher(req *model.Request) ([]byte, error) {
	h := utils.NewHttpSend(req.Url)
	h.SetMethod(req.Method)
	h.SetProxy(req.Proxy)
	h.SetHeader(req.Headers)
	if len(req.Body) > 0 {
		b, _ := json.Marshal(req.Body)
		h.SetBody(string(b))
	}

	return h.Do()
}

func (i twitterPostAnalyzer) Analyze(ctx context.Context, input []byte, task model.Task, req model.Request) (model.AnalyzerResult, error) {
	return model.AnalyzerResult{}, errors.New("twitter post analyze error")
}

func (i twitterPostAnalyzer) Type() analyzer.Type {
	return analyzer.TypeTwitterPost
}
