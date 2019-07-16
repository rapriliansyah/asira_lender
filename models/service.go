package models

type (
	Service struct {
		BaseModel
		Name string `json:"name" gorm:"column:name;type:varchar(255)"`
	}
)
