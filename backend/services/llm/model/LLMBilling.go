// Package model 账单模型
//
// @File: LLMBilling.go
// @Desc: 账单模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	LLMBillingModel interface{}

	LLMBilling struct {
		ID            string           `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                         // ID
		UserID        string           `json:"user_id" gorm:"column:user_id;type:varchar(32);not null"`                 // 用户ID
		ModelMetaID   string           `json:"model_meta_id" gorm:"column:model_meta_id;type:varchar(32);not null"`     // 模型ID
		RequestLogID  string           `json:"request_log_id" gorm:"column:request_log_id;type:varchar(32);not null"`   // 请求日志ID
		InputTokens   int              `json:"input_tokens" gorm:"column:input_tokens;type:int;not null"`               // 输入令牌数
		OutputTokens  int              `json:"output_tokens" gorm:"column:output_tokens;type:int;not null"`             // 输出令牌数
		Cost          float64          `json:"cost" gorm:"column:cost;type:decimal(10,6);not null"`                     // 消耗金额
		Currency      CurrencyTypeEnum `json:"currency" gorm:"column:currency;type:int;not null"`                       // 货币类型
		Status        int              `json:"status" gorm:"column:status;type:int;not null"`                           // 状态[0:未扣费;1:已扣费;2:已退款]
		BalanceBefore float64          `json:"balance_before" gorm:"column:balance_before;type:decimal(12,6);not null"` // 调用前余额
		BalanceAfter  float64          `json:"balance_after" gorm:"column:balance_after;type:decimal(12,6);not null"`   // 调用后余额
		CreatedAt     int64            `json:"created_at" gorm:"autoCreateTime:milli"`                                  // 创建时间
		UpdatedAt     int64            `json:"updated_at" gorm:"autoUpdateTime:milli"`                                  // 更新时间
		DeletedAt     *gorm.DeletedAt  `json:"deleted_at"`                                                              // 删除时间
	}

	defaultLLMBillingModel struct {
		db *gorm.DB
	}
)
