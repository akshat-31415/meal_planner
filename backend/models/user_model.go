package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Email string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time	`json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

