package models

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	ID        int     `gorm:"type:bigint(20)"`
	Wallet_id int     `gorm:"foreignKey:WalletRefer;type:bigint(20)"`
	Name      string  `gorm:"type:varchar(255)"`
	Symbol    string  `gorm:"varchar(100)"`
	Network   string  `gorm:"varchar(100)"`
	Address   string  `gorm:"varchar(42)"`
	Balance   float32 `gorm:"type:decimal(16,8)"`
}
