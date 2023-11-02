package model

type Request struct {
	Url           string
	Proxy         string
	Method        string
	Headers       map[string]string
	Body          string
	Template      string
	Params        string
	ParentRequest ParentRequest
}
type ParentRequest struct {
	Url      string
	Proxy    string
	Method   string
	Headers  map[string]string
	Body     string
	Template string
	Params   string
}

type Item struct {
	Url     string
	Type    string
	Task    Task
	PayLoad string
}

type AnalyzerResult struct {
	Items []Item
	Tasks []Task
}
