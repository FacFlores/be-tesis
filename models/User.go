package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	DNI       string    `gorm:"uniqueIndex;not null" json:dni"`
	Password  string    `gorm:"not null" json:"password"`
	Role      string    `gorm:"type:varchar(255);not null" json:"role"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
