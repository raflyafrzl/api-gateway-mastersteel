package model

type User struct {
	Id       string `json:"id" bson:"_id" gorm:"Id;type:string;primaryKey"`
	Email    string `json:"email" gorm:"email;type:string"`
	Password string `json:"password" gorm:"password;type:string"`
}
