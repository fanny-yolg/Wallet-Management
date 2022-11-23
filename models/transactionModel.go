package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID             int     `gorm:"type:bigint(20)"`
	Src_wallet_id  int     `gorm:"foreignKey:WalletRefer;type:bigint(20)"`
	Src_asset_id   int     `gorm:"foreignKey:AssetRefer;type:bigint(20)"`
	Dest_wallet_id int     `gorm:"foreignKey:WalletRefer;type:bigint(20)"`
	Dest_asset_id  int     `gorm:"foreignKey:AssetRefer;type:bigint(20)"`
	Amount         float32 `gorm:"type:decimal(16,8)"`
	Gas_fee        float32 `gorm:"type:decimal(16,8)"`
	Total          float32 `gorm:"type:decimal(16,8)"`
}
