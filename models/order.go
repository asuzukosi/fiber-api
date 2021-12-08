package models

import "time"

type Order struct {
	Id           uint      `json:"id" gorm:"primary_key, auto"`
	UserRefer    uint      `json:"user_id"`
	User         User      `gorm:"foreignKey:UserRefer"`
	ProductRefer uint      `json:"product_id"`
	Product      Product   `gorm:"foreignKey:ProductRefer"`
	CreatedAt    time.Time `json:"created_at"`
}
