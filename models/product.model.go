package models

type Product struct {
	ID    int    `json:"id" gorm:"type:int;primary_key"`
	Name  string `json:"name" gorm:"type:varchar(100)"`
	Price int    `json:"price" gorm:"type:int"`
}
