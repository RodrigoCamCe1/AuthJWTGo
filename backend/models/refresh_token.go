package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Token      string    `gorm:"unique;not null" json:"token"`
	ExpiryDate time.Time `json:"expiryDate"`
	Revoked    bool      `gorm:"default:false" json:"revoked"`

	// Relación con el Usuario (Foreign Key)
	UserID uint `json:"userId"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
