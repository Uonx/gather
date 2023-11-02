package model

type Project struct {
	Name       string `json:"name"`
	CronStatus bool   `json:"cron_stats"`
	Cron       string `json:"cron"`
	// Status     int    `json:"status"`
	Task []ProjectTask `json:"task" `
}

type ProjectTask struct {
	Url      string `json:"url"`
	Auth     string `json:"auth"`
	Template string `json:"template"`
	Proxy    string `json:"proxy"`
	// Status   int    `json:"status"`
}
