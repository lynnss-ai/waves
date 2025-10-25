// Package model 充值模型
//
// @File: LLMTopUp.go
// @Desc: 充值模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	LLMTopUpModel interface{}

	LLMTopUp struct {
		ID        string           `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`         // ID
		UserID    string           `json:"user_id" gorm:"column:user_id;type:varchar(32);not null"` // 用户ID
		Amount    float64          `json:"amount" gorm:"column:amount;type:decimal(10,6);not null"` // 充值金额
		Currency  CurrencyTypeEnum `json:"currency" gorm:"column:currency;type:int;not null"`       // 货币类型
		Method    TopUpMethodEnum  `json:"method" gorm:"column:method;type:int;not null"`           // 充值方式
		Status    int              `json:"status" gorm:"column:status;type:int;not null"`           // 状态[0:未到账;1:成功;2:失败]
		CreatedAt int64            `json:"created_at" gorm:"autoCreateTime:milli"`                  // 创建时间
		UpdatedAt int64            `json:"updated_at" gorm:"autoUpdateTime:milli"`                  // 更新时间
		DeletedAt *gorm.DeletedAt  `json:"deleted_at"`                                              // 删除时间
	}

	defaultLLMTopUpModel struct {
		db *gorm.DB
	}
)
