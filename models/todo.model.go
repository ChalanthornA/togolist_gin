package models

type TodoModel struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Todo        string `json:"todo"`
	Description string `json:"description"`
	UserRefer   int    `json:"user_refer"`
}
