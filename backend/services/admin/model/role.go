// Package model 角色模型
//
// @File: role.go
// @Desc: 角色模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	RoleModel interface{}

	Role struct {
		ID        string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`    // ID
		RoleName  string          `json:"role_name" gorm:"column:role_name;type:varchar(32)"` // 角色名
		RoleCode  string          `json:"role_code" gorm:"column:role_code;type:varchar(32)"` // 角色编码
		MenuIds   []string        `json:"menu_ids" gorm:"column:menu_ids;type:varchar(32)"`   // 菜单ID列表
		Status    int             `json:"status" gorm:"column:status;type:int;not null"`      // 状态[0:禁用,1:启用]
		CreatedAt int64           `json:"created_at" gorm:"autoCreateTime:milli"`             // 创建时间
		UpdatedAt int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`             // 更新时间
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`                                         // 删除时间
	}

	defaultRoleModel struct {
		db *gorm.DB
	}
)
