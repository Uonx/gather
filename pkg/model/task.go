package model

type Task struct {
	Id         int64
	TaskId     int64
	Auth       string
	Proxy      string
	Url        string
	Template   string
	Method     string
	Body       []byte
	Params     string
	ParentTask ParentTask
}

type ParentTask struct {
	Id       int64
	TaskId   int64
	Auth     string
	Proxy    string
	Url      string
	Template string
	Method   string
}
