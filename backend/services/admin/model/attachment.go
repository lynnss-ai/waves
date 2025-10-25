// Package model 附件模型
//
// @File: attachment.go
// @Desc: 附件模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	AttachmentModel interface{}

	Attachment struct {
		ID                   string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                                       // ID
		AttachmentCategoryID string          `json:"attachment_category_id" gorm:"column:attachment_category_id;type:varchar(32);not null"` // 分类ID
		FileName             string          `json:"file_name" gorm:"column:file_name;type:varchar(255);not null"`                          // 文件名
		FileSize             int64           `json:"file_size" gorm:"column:file_size;type:bigint"`                                         // 文件大小
		FileURL              string          `json:"file_url" gorm:"column:file_url;type:varchar(255);not null"`                            // 文件URL
		FileType             string          `json:"file_type" gorm:"column:file_type;type:varchar(64)"`                                    // 文件类型
		CreatedAt            int64           `json:"created_at" gorm:"autoCreateTime:milli"`                                                // 创建时间
		UpdatedAt            int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                                                // 更新时间
		DeletedAt            *gorm.DeletedAt `json:"deleted_at"`                                                                            // 删除时间
	}

	defaultAttachmentModel struct {
		db *gorm.DB
	}
)
