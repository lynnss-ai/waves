// Package model LLM模型
//
// @File: llmmeta.go
// @Desc: LLM模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	LLMMetaModel interface{}

	LLMMeta struct {
		ID                  string           `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                                  // ID
		VendorID            string           `json:"vendor_id" gorm:"column:vendor_id;type:varchar(32);not null"`                      // 供应商ID
		ModelName           string           `json:"model_name" gorm:"column:model_name;type:varchar(32);not null"`                    // 模型名
		ModelCode           string           `json:"model_code" gorm:"column:model_code;type:varchar(32);not null"`                    // 模型编码
		ModelVersion        string           `json:"model_version" gorm:"column:model_version;type:varchar(128);not null"`             // 模型版本
		ModelDesc           string           `json:"model_desc" gorm:"column:model_desc;type:varchar(256)"`                            // 模型描述
		ApiUrl              string           `json:"api_url" gorm:"column:api_url;type:varchar(256);not null"`                         // API地址
		ApiKey              string           `json:"api_key" gorm:"column:api_key;type:varchar(256);not null"`                         // API密钥
		IsSupportsEmbedding bool             `json:"is_supports_embedding" gorm:"column:is_supports_embedding;type:bool;not null"`     // 是否支持embedding
		IsSupportsStreaming bool             `json:"is_supports_streaming" gorm:"column:is_supports_streaming;type:bool;not null"`     // 是否支持流式输出
		Extra               string           `json:"extra" gorm:"column:extra;type:text"`                                              // 额外信息
		Permission          int              `json:"permission" gorm:"column:permission;type:int;not null"`                            // 权限[1:公开;2:私有]
		PricePerInputToken  float64          `json:"price_per_input_token" gorm:"column:price_per_input_token;type:float;not null"`    // 输入token价格
		PricePerOutputToken float64          `json:"price_per_output_token" gorm:"column:price_per_output_token;type:float;not null"`  // 输出token价格
		ContextLength       int64            `json:"context_length" gorm:"column:context_length;type:bigint;not null"`                 // 上下文长度
		MaxTokensPerRequest int64            `json:"max_tokens_per_request" gorm:"column:max_tokens_per_request;type:bigint;not null"` // 最大token数
		CurrencyType        CurrencyTypeEnum `json:"currency_type" gorm:"column:currency_type;type:int;not null"`                      // 货币类型
		CreatedAt           int64            `json:"created_at" gorm:"autoCreateTime:milli"`                                           // 创建时间
		UpdatedAt           int64            `json:"updated_at" gorm:"autoUpdateTime:milli"`                                           // 更新时间
		DeletedAt           *gorm.DeletedAt  `json:"deleted_at"`                                                                       // 删除时间
	}

	defaultLLMMetaModel struct {
		db *gorm.DB
	}
)
