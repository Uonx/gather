package entities

type Project struct {
	Model
	Name string `json:"name"`
	Type int    `json:"type"`
	//优先级
	//
	CronStatus bool   `json:"cron_stats"`
	Cron       string `json:"cron"`
	Status     int    `json:"status"`
	Task       []Task `json:"task" gorm:"foreignKey:ProjectId;references:Id"`
}
