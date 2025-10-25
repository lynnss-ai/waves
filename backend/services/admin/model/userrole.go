// Package model 用户角色模型
//
// @File: userrole.go
// @Desc: 用户角色模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	UserRoleModel interface {
	}

	UserRole struct {
		ID        string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"` // ID
		UserID    string          `json:"user_id" gorm:"column:user_id;type:varchar(32)"`  // 用户ID
		RoleID    string          `json:"role_id" gorm:"column:role_id;type:varchar(32)"`  // 角色ID
		CreatedAt int64           `json:"created_at" gorm:"autoCreateTime:milli"`          // 创建时间
		UpdatedAt int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`          // 更新时间
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`                                      // 删除时间
	}

	defaultUserRoleModel struct {
		db *gorm.DB
	}
)
