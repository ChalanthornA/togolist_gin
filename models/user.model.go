package models

type UserModel struct {
	ID       int  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
