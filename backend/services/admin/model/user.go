// Package model 用户模型
//
// @File: user.go
// @Desc: 用户模型
// @Date: 2025-10-25

package model

import (
	"gorm.io/gorm"
)

type (
	UserModel interface {
	}

	User struct {
		ID          string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`             // ID
		Name        string          `json:"name" gorm:"column:name;type:varchar(32)"`                    // 姓名
		UserName    string          `json:"user_name" gorm:"column:user_name;type:varchar(32);not null"` // 用户名
		Email       string          `json:"email" gorm:"column:email;type:varchar(255);not null"`        // 邮箱
		Password    string          `json:"password" gorm:"column:password;type:varchar(255)"`           // 密码
		Salt        string          `json:"salt" gorm:"column:salt;type:varchar(64)"`                    // 盐值
		IsChangePwd bool            `json:"is_change_pwd" gorm:"column:is_change_pwd;type:bool"`         // 是否修改密码[true:是,false:否]
		LastLoginAt int64           `json:"last_login_at" gorm:"column:last_login_at;type:bigint"`       // 最后登录时间
		CreatedAt   int64           `json:"created_at" gorm:"autoCreateTime:milli"`                      // 创建时间
		UpdatedAt   int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                      // 更新时间
		DeletedAt   *gorm.DeletedAt `json:"deleted_at"`                                                  // 删除时间
	}

	defaultUserModel struct {
		db *gorm.DB
	}
)
