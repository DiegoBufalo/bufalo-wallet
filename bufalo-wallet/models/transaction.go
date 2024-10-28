package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount     float64   `json:"amount" gorm:"not null"`
	Date       time.Time `json:"date" gorm:"not null"`
	Type       string    `json:"type" gorm:"type:varchar(50);not null"` // Tipo pode ser "entrada" ou "saída"
	UserID     uint      `json:"user_id" gorm:"not null"`               // Chave estrangeira para User
	User       User      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryID uint      `json:"category_id" gorm:"not null"` // Chave estrangeira para Category
	Category   Category  `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Recurring  bool      `json:"recurring" gorm:"not null"`
	PeriodDays int       `json:"period_days" gorm:"not null"`
}
