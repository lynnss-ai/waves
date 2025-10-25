// Package model 附件分类模型
//
// @File: attachmentcategory.go
// @Desc: 附件分类模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	AttachmentCategoryModel interface{}

	AttachmentCategory struct {
		ID           string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                     // ID
		CategoryName string          `json:"category_name" gorm:"column:category_name;type:varchar(32);not null"` // 分类名
		CategoryCode string          `json:"category_code" gorm:"column:category_code;type:varchar(32);not null"` // 分类编码
		Order        int             `json:"order" gorm:"column:order;type:int;not null"`                         // 排序
		CreatedAt    int64           `json:"created_at" gorm:"autoCreateTime:milli"`                              // 创建时间
		UpdatedAt    int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                              // 更新时间
		DeletedAt    *gorm.DeletedAt `json:"deleted_at"`                                                          // 删除时间
	}

	defaultAttachmentCategoryModel struct {
		db *gorm.DB
	}
)
