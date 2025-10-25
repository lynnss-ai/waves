// Package model 菜单模型
//
// @File: menu.go
// @Desc: 菜单模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

const (
	MenuTypeDir  MenuTypeEnum = iota + 1 // 目录
	MenuTypeMenu                         // 菜单
	MenuTypeBtn                          // 按钮
)

type (
	MenuModel interface{}

	MenuTypeEnum int

	Menu struct {
		ID                   string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                           // ID
		ParentID             string          `json:"parent_id" gorm:"column:parent_id;type:varchar(32)"`                        // 父菜单ID
		MenuName             string          `json:"menu_name" gorm:"column:menu_name;type:varchar(32);not null"`               // 菜单名
		MenuCode             string          `json:"menu_code" gorm:"column:menu_code;type:varchar(32);not null"`               // 菜单编码
		MenuPath             string          `json:"menu_path" gorm:"column:menu_path;type:varchar(255)"`                       // 菜单路径
		MenuComponent        string          `json:"menu_component" gorm:"column:menu_component;type:varchar(255)"`             // 菜单组件
		MenuIcon             string          `json:"menu_icon" gorm:"column:menu_icon;type:varchar(255)"`                       // 菜单图标
		MenuType             MenuTypeEnum    `json:"menu_type" gorm:"column:menu_type;type:int;not null"`                       // 菜单类型[0:目录,1:菜单,2:按钮]
		IsHideInMenu         bool            `json:"is_hide_in_menu" gorm:"column:is_hide_in_menu;type:bool"`                   // 是否在菜单中隐藏[true:隐藏,false:显示]
		IsHideChildrenInMenu bool            `json:"is_hide_children_in_menu" gorm:"column:is_hide_children_in_menu;type:bool"` // 是否在菜单中隐藏子菜单[true:隐藏,false:显示]
		IsHideInBreadcrumb   bool            `json:"is_hide_in_breadcrumb" gorm:"column:is_hide_in_breadcrumb;type:bool"`       // 是否在面包屑中隐藏[true:隐藏,false:显示]
		Status               int             `json:"status" gorm:"column:status;type:int;not null"`                             // 状态[0:禁用,1:启用]
		Order                int             `json:"order" gorm:"column:order;type:int;not null"`                               // 排序
		CreatedAt            int64           `json:"created_at" gorm:"autoCreateTime:milli"`                                    // 创建时间
		UpdatedAt            int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                                    // 更新时间
		DeletedAt            *gorm.DeletedAt `json:"deleted_at"`                                                                // 删除时间
	}

	TreeMenu struct {
		Menu
		Children []*TreeMenu
	}

	defaultMenuModel struct {
		db *gorm.DB
	}
)
