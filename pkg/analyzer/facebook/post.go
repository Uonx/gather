package facebook

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Uonx/gather/pkg/analyzer"
	"github.com/Uonx/gather/pkg/model"
	"github.com/Uonx/gather/pkg/utils"
)

func init() {
	analyzer.RegistryAnalyzer(&facebookUserAboutAnalyzer{})
}

type facebookUserAboutAnalyzer struct{}

var headerParam = regexp.MustCompile(`("token":\s*"[^"]+")|("rev":\s*[^}]+)|("ACCOUNT_ID":\s*[^,]+)|("haste_session":\s*"[^"]+")|("e":\s*"[^"]+")|("selectedID":\s*"[^"]+")`)
var numberData = regexp.MustCompile(`},([0-9]+)]`)
var cookie = `sb=66goZRERghG58d0R0hrqpygs; datr=66goZQfmUkI7Cr8FvQnb8Twa; locale=zh_CN; wl_cbv=v2%3Bclient_version%3A2346%3Btimestamp%3A1698726740; vpd=v1%3B685x400x2; c_user=100025702010490; xs=36%3AWRyTk9L9UBvZ4A%3A2%3A1698808060%3A-1%3A-1; fr=1bzLq59dTcC5G2RFC.AWWtxS5jRBqNUbVioyAOIDxKwSQ.BlQaVP.XI.AAA.0.0.BlQcEA.AWV7mZDlIJU;`

func (i facebookUserAboutAnalyzer) SetHeader(_ string, req *model.Request) (map[string]string, error) {
	headers := make(map[string]string)
	headers["authority"] = "www.facebook.com"
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
	headers["accept-language"] = "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6"
	headers["Sec-Ch-Prefers-Color-Scheme"] = `light`
	headers["Sec-Ch-Ua"] = `"Chromium";v="118", "Microsoft Edge";v="118", "Not=A?Brand";v="99"`
	headers["sec-ch-ua-full-version-list"] = `"Chromium";v="118.0.5993.118", "Microsoft Edge";v="118.0.2088.76", "Not=A?Brand";v="99.0.0.0"`

	headers["Sec-Ch-Ua-Mobile"] = `?0`
	headers["Sec-Ch-Ua-Platform"] = `"Windows"`
	headers["Sec-Ch-Ua-Platform-Version"] = `"15.0.0"`
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.76"
	headers["viewport-width"] = "1392"
	headers["Cache-Control"] = "no-cache"
	headers["Cookie"] = cookie
	h := utils.NewHttpSend("https://www.facebook.com")
	h.SetProxy(req.Proxy)
	h.SetHeader(headers)
	header_param_by, err := h.Do()
	if err != nil {
		return nil, err
	}
	params := headerParam.FindAllSubmatch(header_param_by, -1)
	token := ""
	lsdtoken := ""
	rev := ""
	userID := ""
	hasteSession := ""
	hsi := ""
	variable_id := ""
	for _, v := range params {
		ps := strings.Split(string(v[0]), ":")
		if strings.Trim(ps[0], "\"") == "token" {
			s := strings.Split(string(v[0]), ":")
			if len(s) == 2 {
				lsdtoken = strings.Trim(strings.Split(string(v[0]), ":")[1], "\"")
			} else {
				token = strings.Trim(strings.Join(s[1:], ":"), "\"")
			}

		}
		if strings.Trim(ps[0], "\"") == "rev" {
			rev = strings.Trim(strings.Split(string(v[0]), ":")[1], "\"")
		}
		if strings.Trim(ps[0], "\"") == "selectedID" {
			variable_id = strings.Trim(strings.Split(string(v[0]), ":")[1], "\"")
		}
		if strings.Trim(ps[0], "\"") == "e" {
			hsi = strings.Trim(strings.Split(string(v[0]), ":")[1], "\"")
		}
		if strings.Trim(ps[0], "\"") == "ACCOUNT_ID" {
			userID = strings.Trim(strings.Split(string(v[0]), ":")[1], "\"")
		}
		if strings.Trim(ps[0], "\"") == "haste_session" {
			s := strings.Split(string(v[0]), ":")
			if len(s) == 3 {
				hasteSession = strings.Trim(strings.Join(s[1:], ":"), "\"")
			}
		}
	}
	variable_id = "100044328734711"
	headers["X-Asbd-Id"] = "129477"
	headers["X-Fb-Friendly-Name"] = "CometNotificationsDropdownQuery"
	headers["X-Fb-Lsd"] = lsdtoken
	headers["fb_dtsg"] = token
	headers["rev"] = rev
	headers["user_id"] = userID
	headers["haste_session"] = hasteSession
	headers["hsi"] = hsi
	headers["variable_id"] = variable_id
	return headers, nil
}

func (i facebookUserAboutAnalyzer) Fetcher(req *model.Request) ([]byte, error) {
	h := utils.NewHttpSend(req.Url)
	if len(req.Method) > 0 {
		h.SetMethod(req.Method)
	}

	h.SetMethod(http.MethodPost)
	header := map[string]string{}
	header["User-Agent"] = "Fiddler"
	header["Content-Type"] = "application/x-www-form-urlencoded"
	header["X-Fb-Friendly-Name"] = "ProfileCometTimelineFeedRefetchQuery"
	header["X-Fb-Lsd"] = req.Headers["X-Fb-Lsd"]
	header["X-Asbd-Id"] = "129477"
	header["Cookie"] = cookie
	header["Host"] = "www.facebook.com"
	header["Content-Length"] = "2382"
	h.SetProxy(req.Proxy)
	h.SetHeader(header)
	return h.Do()
}

func (i facebookUserAboutAnalyzer) Analyze(ctx context.Context, input []byte, task model.Task, req model.Request) (model.AnalyzerResult, error) {
	result := model.AnalyzerResult{}
	s := strings.Split(string(input), "\n")
	for k, v := range s {
		if k == 0 {
			data := DataEx{}
			err := json.Unmarshal([]byte(v), &data)
			if err != nil {
				continue
			}

			result.Items = append(result.Items, model.Item{
				Url:  req.Url,
				Type: "facebook_post",
				Task: task,
				PayLoad: fmt.Sprintf("message: %s like_cout: %s total_count: %s share_count: %d \n",
					data.Data.Node.TimelineListFeedUnits.Edges[0].Node.CometSections.Content.Story.CometSections.Message.Story.Message.Text,
					data.Data.Node.TimelineListFeedUnits.Edges[0].Node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.UfiRenderer.Facebook.CometUfiSummaryAndActionsRenderer.Facebook.I18nReactionCount,
					data.Data.Node.TimelineListFeedUnits.Edges[0].Node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.UfiRenderer.Facebook.CometUfiSummaryAndActionsRenderer.Facebook.I18nShareCount,
					data.Data.Node.TimelineListFeedUnits.Edges[0].Node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.UfiRenderer.Facebook.CometUfiSummaryAndActionsRenderer.Facebook.TotalCommentCount,
				),
			})
		} else {
			data := EdgeNode{}
			err := json.Unmarshal([]byte(v), &data)
			if err != nil {
				continue
			}

			if data.Label != "ProfileCometTimelineFeed_user$stream$ProfileCometTimelineFeed_user_timeline_list_feed_units" {
				continue
			}
			result.Items = append(result.Items, model.Item{
				Url:  req.Url,
				Type: "facebook_post",
				Task: task,
				PayLoad: fmt.Sprintf("message: %s like_cout: %s total_count: %s share_count: %d \n",
					data.Node.Node.CometSections.Content.Story.CometSections.Message.Story.Message.Text,
					data.Node.Node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.UfiRenderer.Facebook.CometUfiSummaryAndActionsRenderer.Facebook.I18nReactionCount,
					data.Node.Node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.UfiRenderer.Facebook.CometUfiSummaryAndActionsRenderer.Facebook.I18nShareCount,
					data.Node.Node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.UfiRenderer.Facebook.CometUfiSummaryAndActionsRenderer.Facebook.TotalCommentCount,
				),
			})

		}
	}

	reg := regexp.MustCompile(`"end_cursor":(\s*[^,]+)`)
	cursor := reg.Find(input)
	reg_cursor := regexp.MustCompile(`"cursor":(\s*[^,]+)`)

	url := (reg_cursor.ReplaceAllString(req.Url, fmt.Sprintf(`"cursor":%s`, cursor)))
	result.Tasks = append(result.Tasks, model.Task{
		TaskId:   task.TaskId,
		Auth:     task.Auth,
		Proxy:    task.Proxy,
		Url:      url,
		Method:   task.Method,
		Template: string(analyzer.TypeFaceBookPost),
	})
	return result, nil
}

func (i facebookUserAboutAnalyzer) Type() analyzer.Type {
	return analyzer.TypeFaceBookPost
}

type DataEx struct {
	Data Data `json:"data"`
}

type Data struct {
	Node Node `json:"node"`
}
type Node struct {
	Id                    string                `json:"id"`
	TimelineListFeedUnits TimelineListFeedUnits `json:"timeline_list_feed_units"`
}

type TimelineListFeedUnits struct {
	Edges []Edge `json:"edges"`
}

type EdgeNode struct {
	Label string `json:"label"`
	Node  Edge   `json:"data"`
}

type Edge struct {
	Node   NodeInfo `json:"node"`
	Cursor string   `json:"cursor"`
}

type NodeInfo struct {
	Id            string         `json:"id"`
	PostId        string         `json:"post_id"`
	CometSections CometSectionsZ `json:"comet_sections"`
}

type CometSectionsZ struct {
	Content  Content   `json:"content"`
	Feedback FeedbackX `json:"feedback"`
}

type Content struct {
	Story StoryY `json:"story"`
}

type StoryY struct {
	Id            string        `json:"id"`
	CometSections CometSections `json:"comet_sections"`
}

type CometSections struct {
	Message MessageZ `json:"message"`
}

type MessageZ struct {
	Story StoryZ `json:"story"`
}
type StoryZ struct {
	Id      string  `json:"id"`
	Message Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
}

type FeedbackX struct {
	Story Story `json:"story"`
}

type Story struct {
	FeedbackContext FeedbackContext `json:"feedback_context"`
}
type FeedbackContext struct {
	FeedbackTargetWithContext FeedbackTargetWithContext `json:"feedback_target_with_context"`
}

type FeedbackTargetWithContext struct {
	UfiRenderer UfiRenderer `json:"ufi_renderer"`
}

type UfiRenderer struct {
	Facebook FeedbackY `json:"feedback"`
}

type FeedbackY struct {
	Id                                string                            `json:"id"`
	CometUfiSummaryAndActionsRenderer CometUfiSummaryAndActionsRenderer `json:"comet_ufi_summary_and_actions_renderer"`
	Url                               string                            `json:"url"`
}

type CometUfiSummaryAndActionsRenderer struct {
	Facebook FeedbackZ `json:"feedback"`
}

type FeedbackZ struct {
	I18nShareCount    string `json:"i18n_share_count"`
	TotalCommentCount int    `json:"total_comment_count"`
	I18nReactionCount string `json:"i18n_reaction_count"`
}
