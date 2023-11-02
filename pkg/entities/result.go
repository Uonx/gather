package entities

type Result struct {
	Model
	TaskId  int64  `json:"task_id"`
	SubId   int64  `json:"sub_id"`
	PayLoad string `json:"pay_load"`
}
