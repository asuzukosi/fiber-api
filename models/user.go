package models

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey, auto"`
	CreatedAt time.Time `json:"createdAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}
