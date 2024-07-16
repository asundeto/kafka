package internal

import "time"

type Message struct {
	Id      uint      `gorm:"primaryKey" json:"id"`
	Message string    `gorm:"not null" json:"message"`
	Created time.Time `gorm:"not null" json:"created"`
}
