// Package model 文章分类模型
//
// @File: articlecategory.go
// @Desc: 文章分类模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	ArticleCategoryModel interface{}

	ArticleCategory struct {
		ID           string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                                 // ID
		ParentID     string          `json:"parent_id" gorm:"column:parent_id;type:varchar(32)"`                              // 父分类ID
		CategoryName string          `json:"category_name" gorm:"column:category_name;type:varchar(64);not null"`             // 分类名
		CategoryCode string          `json:"category_code" gorm:"column:category_code;type:varchar(64);not null;uniqueIndex"` // 分类编码
		Slug         string          `json:"slug" gorm:"column:slug;type:varchar(512)"`                                       // URL别名
		CategoryDesc string          `json:"category_desc" gorm:"column:category_desc;type:varchar(255)"`                     // 分类描述
		Order        int             `json:"order" gorm:"column:order;type:int"`                                              // 排序
		Status       int             `json:"status" gorm:"column:status;type:tinyint(1)"`                                     // 状态[0:禁用,1:启用]
		CreatedAt    int64           `json:"created_at" gorm:"autoCreateTime:milli"`                                          // 创建时间
		UpdatedAt    int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                                          // 更新时间
		DeletedAt    *gorm.DeletedAt `json:"deleted_at"`                                                                      // 删除时间
	}

	defaultArticleCategoryModel struct {
		db *gorm.DB
	}
)
