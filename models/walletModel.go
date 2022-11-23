package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	ID   int    `gorm:"type:bigint(20)"`
	Name string `gorm:"type:varchar(200)"`
}
