package model

type Employee struct {
	Id   string `json:"id" gorm:"primaryKey;type:string;id"`
	Name string `json:"name" gorm:"type:string;name"`
}
