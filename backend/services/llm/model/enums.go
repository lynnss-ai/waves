// Package model 枚举模型
//
// @File: enums.go
// @Desc: 枚举模型
// @Date: 2025-10-25
package model

// CurrencyTypeEnum 货币类型枚举
type CurrencyTypeEnum int

const (
	CurrencyTypeCNY CurrencyTypeEnum = iota + 1 // 人民币
	CurrencyTypeUSD                             // 美元
)

// 充值方式枚举
type TopUpMethodEnum int

const (
	TopUpMethodAlipay TopUpMethodEnum = iota + 1 // 支付宝
	TopUpMethodWechat                            // 微信
)
