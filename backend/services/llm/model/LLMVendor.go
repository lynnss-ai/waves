// Package model LLM供应商模型
//
// @File: llmvendor.go
// @Desc: LLM供应商模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	LLMVendorModel interface{}

	LLMVendor struct {
		ID         string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                  // ID
		VendorName string          `json:"vendor_name" gorm:"column:vendor_name;type:varchar(32);not null"`  // 供应商名
		VendorCode string          `json:"vendor_code" gorm:"column:vendor_code;type:varchar(32);not null"`  // 供应商编码
		VendorIcon string          `json:"vendor_icon" gorm:"column:vendor_icon;type:varchar(256);not null"` // 供应商图标
		WebSite    string          `json:"web_site" gorm:"column:web_site;type:varchar(256);not null"`       // 供应商网站
		Status     int             `json:"status" gorm:"column:status;type:int;not null"`                    // 状态[1:启用;0:禁用]
		Order      int             `json:"order" gorm:"column:order;type:int;not null"`                      // 排序
		CreatedAt  int64           `json:"created_at" gorm:"autoCreateTime:milli"`                           // 创建时间
		UpdatedAt  int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                           // 更新时间
		DeletedAt  *gorm.DeletedAt `json:"deleted_at"`                                                       // 删除时间
	}

	defaultLLMVendorModel struct {
		db *gorm.DB
	}
)
