// Package model LLM用户实时余额模型
//
// @File: LLMUserBalance.go
// @Desc: LLM用户实时余额模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	LLMUserBalanceModel interface{}

	LLMUserBalance struct {
		ID        string           `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`           // ID
		UserID    string           `json:"user_id" gorm:"column:user_id;type:varchar(32);not null"`   // 用户ID
		Balance   float64          `json:"balance" gorm:"column:balance;type:decimal(10,6);not null"` // 余额
		Currency  CurrencyTypeEnum `json:"currency" gorm:"column:currency;type:int;not null"`         // 货币类型
		CreatedAt int64            `json:"created_at" gorm:"autoCreateTime:milli"`                    // 创建时间
		UpdatedAt int64            `json:"updated_at" gorm:"autoUpdateTime:milli"`                    // 更新时间
		DeletedAt *gorm.DeletedAt  `json:"deleted_at"`                                                // 删除时间
	}

	defaultLLMUserBalanceModel struct {
		db *gorm.DB
	}
)
