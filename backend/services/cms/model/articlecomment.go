// Package model 文章评论模型
//
// @File: articlecomment.go
// @Desc: 文章评论模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	ArticleCommentModel interface{}

	ArticleComment struct {
		ID        string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                 // ID
		ArticleID string          `json:"article_id" gorm:"column:article_id;type:varchar(32);not null"`   // 文章ID
		UserID    string          `json:"user_id" gorm:"column:user_id;type:varchar(32);not null"`         // 用户ID
		ParentID  string          `json:"parent_id" gorm:"column:parent_id;type:varchar(32);not null"`     // 父评论ID[支持多级平论]
		ReplyToID string          `json:"reply_to_id" gorm:"column:reply_to_id;type:varchar(32);not null"` // 被回复的评论ID（精确到子层）
		Content   string          `json:"content" gorm:"column:content;type:text"`                         // 评论内容
		LikeCount int             `json:"like_count" gorm:"column:like_count;type:int;default:0"`          // 点赞量
		Status    int             `json:"status" gorm:"column:status;type:int;default:0"`                  // 状态(1:正常,0:隐藏/屏蔽)
		IP        string          `json:"ip" gorm:"column:ip;type:varchar(64)"`                            // 评论IP
		UserAgent string          `json:"user_agent" gorm:"column:user_agent;type:varchar(255)"`           // 设备信息
		CreatedAt int64           `json:"created_at" gorm:"autoCreateTime:milli"`                          // 创建时间
		UpdatedAt int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                          // 更新时间
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`                                                      // 删除时间
	}

	defaultArticleCommentModel struct {
		db *gorm.DB
	}
)
