// Package model 权限模型
//
// @File: permission.go
// @Desc: 权限模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

const (
	PermissionTypeMenu PermissionTypeEnum = iota + 1 // 菜单
	PermissionTypeApi                                // 接口
)

type (
	PermissionModel interface{}

	PermissionTypeEnum int

	Permission struct {
		ID             string             `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                         // ID
		PermissionName string             `json:"permission_name" gorm:"column:permission_name;type:varchar(64);not null"` // 权限名称
		PermissionCode string             `json:"permission_code" gorm:"column:permission_code;type:varchar(64);not null"` // 权限编码
		PermissionDesc string             `json:"permission_desc" gorm:"column:permission_desc;type:varchar(256)"`         // 权限描述
		PermissionType PermissionTypeEnum `json:"permission_type" gorm:"column:permission_type;type:int;not null"`         // 权限类型
		Method         string             `json:"method" gorm:"column:method;type:varchar(16)"`                            // HTTP方法[GET\POST\PUT\DELETE]API类型必填
		Path           string             `json:"path" gorm:"column:path;type:varchar(256)"`                               // 接口路径
		Action         string             `json:"action" gorm:"column:action;type:varchar(64)"`                            // 操作描述
		MenuId         string             `json:"menu_id" gorm:"column:menu_id;type:varchar(32)"`                          // 菜单ID
		Status         int                `json:"status" gorm:"column:status;type:int;not null"`                           // 状态[1:启用;0:禁用]
		Extra          string             `json:"extra" gorm:"column:extra;type:text"`                                     // 扩展信息[允许的HTTP Query白名单、限制等]
		CreatedAt      int64              `json:"created_at" gorm:"autoCreateTime:milli"`                                  // 创建时间
		UpdatedAt      int64              `json:"updated_at" gorm:"autoUpdateTime:milli"`                                  // 更新时间
		DeletedAt      *gorm.DeletedAt    `json:"deleted_at"`                                                              // 删除时间
	}

	defaultPermissionModel struct {
		db *gorm.DB
	}
)
