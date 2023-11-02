package entities

type Task struct {
	Model
	ProjectId int       `json:"project_id"`
	Url       string    `json:"url"`
	Auth      string    `json:"auth"`
	Template  string    `json:"template"`
	Proxy     string    `json:"proxy"`
	Status    int       `json:"status"`
	SubTask   []SubTask `gorm:"foreignKey:TaskId;references:Id"`
}
