package models

type Product struct {
	Id         int    `json:"id" gorm:"column:ProductId;primaryKey;autoIncrement"`
	Name       string `json:"name" gorm:"column:Name;not null"`
	Ingredient string `json:"ingredient" gorm:"column:Ingredient;not null"`
	CategoryId int    `json:"categoryId" gorm:"column:CategoryId;not null"`
}
