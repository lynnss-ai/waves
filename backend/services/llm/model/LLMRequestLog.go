// Package model LLM请求日志模型
//
// @File: llmrequestlog.go
// @Desc: LLM请求日志模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	LLMRequestLogModel interface{}

	LLMRequestLog struct {
		ID             string  `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                 // ID
		LLMMetaID      string  `json:"llm_meta_id" gorm:"column:llm_meta_id;type:varchar(32);not null"` // LLM模型ID
		UserID         string  `json:"user_id" gorm:"column:user_id;type:varchar(32);not null"`         // 用户ID
		RequestPrompt  string  `json:"request_prompt" gorm:"column:request_prompt;type:text"`           // 请求内容
		ResponseResult string  `json:"response_result" gorm:"column:response_result;type:text"`         // 响应内容
		InputTokens    int64   `json:"input_tokens" gorm:"column:input_tokens;type:bigint;not null"`    // 输入token数
		OutputTokens   int64   `json:"output_tokens" gorm:"column:output_tokens;type:bigint;not null"`  // 输出token数
		Cost           float64 `json:"cost" gorm:"column:cost;type:decimal(10,6);not null"`             // 本次调用费用
		Status         int     `json:"status" gorm:"column:status;type:int;not null"`                   // 状态[1:成功;0:失败]
		ErrorMsg       string  `json:"error_msg" gorm:"column:error_msg;type:text"`                     // 错误信息
		TotalTokens    int64   `json:"total_tokens" gorm:"column:total_tokens;type:bigint;not null"`    // 总token数
	}

	defaultLLMRequestLogModel struct {
		db *gorm.DB
	}
)
