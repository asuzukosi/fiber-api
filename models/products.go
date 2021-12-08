package models

import "time"

type Product struct {
	Id           uint      `json:"id" gorm:"primaryKey, auto"`
	CreatedAt    time.Time `json:"created"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serialNumber"`
}
