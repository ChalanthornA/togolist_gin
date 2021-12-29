package models

type TodoModel struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Todo        string `json:"todo"`
	Description string `json:"description"`
	UserRefer   int64
	User        UserModel `json:"user" gorm:"foreignKey:user_refer;references:id"`
}
