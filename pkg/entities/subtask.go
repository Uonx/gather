package entities

type SubTask struct {
	Model
	TaskId     int    `json:"task_id"`
	Url        string `json:"url"`
	Auth       string `json:"auth"`
	Template   string `json:"template"`
	Proxy      string `json:"proxy"`
	Params     string `json:"params"`
	Status     int    `json:"status"`
	ErrMessage string `json:"err_message"`
	Task       Task   `gorm:"foreignKey:Id;references:TaskId"`
}
