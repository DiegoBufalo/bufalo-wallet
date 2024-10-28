package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(100);not null"`
	Type string `json:"type" gorm:"type:varchar(50);not null"` // Tipo pode ser "entrada" ou "saída"
}
