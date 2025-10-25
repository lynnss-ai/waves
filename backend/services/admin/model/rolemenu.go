// Package model 角色菜单模型
//
// @File: rolemenu.go
// @Desc: 角色菜单模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	RoleMenuModel interface{}

	RoleMenu struct {
		ID        string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`         // ID
		RoleId    string          `json:"role_id" gorm:"column:role_id;type:varchar(32);not null"` // 角色ID
		MenuId    string          `json:"menu_id" gorm:"column:menu_id;type:varchar(32);not null"` // 菜单ID
		CreatedAt int64           `json:"created_at" gorm:"autoCreateTime:milli"`                  // 创建时间
		UpdatedAt int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                  // 更新时间
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`                                              // 删除时间
	}

	defaultRoleMenuModel struct {
		db *gorm.DB
	}
)
