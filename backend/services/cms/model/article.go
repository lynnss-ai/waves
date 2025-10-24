// Package model 文章模型
//
// @File: article.go
// @Desc: 文章模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

// 文章模型
type (
	ArticleModel interface{}

	Article struct {
		ID           string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                 // ID
		CategoryID   string          `json:"category_id" gorm:"column:category_id;type:varchar(32);not null"` // 分类ID
		Title        string          `json:"title" gorm:"column:title;type:varchar(255);not null"`            // 标题
		Slug         string          `json:"slug" gorm:"column:slug;type:varchar(512)"`                       // URL别名
		Excerpt      string          `json:"excerpt" gorm:"column:excerpt;type:varchar(255);not null"`        // 摘要
		Thumbnail    string          `json:"thumbnail" gorm:"column:thumbnail;type:varchar(255)"`             // 封面图
		Content      string          `json:"content" gorm:"column:content;type:text"`                         // 内容
		Tags         string          `json:"tags" gorm:"column:tags;type:varchar(512)"`                       // 标签[用逗号分隔]
		Status       int             `json:"status" gorm:"column:status;type:int;default:0"`                  // 状态(0:草稿,1:已发布,2:已下线)
		AuthorID     string          `json:"author_id" gorm:"column:author_id;type:varchar(32);not null"`     // 作者ID
		ViewCount    int             `json:"view_count" gorm:"column:view_count;type:int"`                    // 访问量
		CommentCount int             `json:"comment_count" gorm:"column:comment_count;type:int"`              // 评论量
		LikeCount    int             `json:"like_count" gorm:"column:like_count;type:int"`                    // 点赞量
		PublishedAt  int64           `json:"published_at" gorm:"column:published_at;type:bigint"`             // 发布时间
		CreatedAt    int64           `json:"created_at" gorm:"autoCreateTime:milli"`                          // 创建时间
		UpdatedAt    int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                          // 更新时间
		DeletedAt    *gorm.DeletedAt `json:"deleted_at"`                                                      // 删除时间
	}
)
