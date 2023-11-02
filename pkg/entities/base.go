package entities

type Model struct {
	Id int `json:"pkid" gorm:"column:pkid;primaryKey;autoIncrement;comment:主键编码"`
}
