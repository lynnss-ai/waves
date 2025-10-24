// Package model 字典模型
//
// @File: dictionary.go
// @Desc: 字典模型
// @Date: 2025-10-25

package model

import "gorm.io/gorm"

type (
	DictionaryModel interface{}

	Dictionary struct {
		ID         string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`             // ID
		ParentID   string          `json:"parent_id" gorm:"column:parent_id;type:varchar(32)"`          // 父字典ID
		DicName    string          `json:"dic_name" gorm:"column:dic_name;type:varchar(32);not null"`   // 字典名
		DicCode    string          `json:"dic_code" gorm:"column:dic_code;type:varchar(32);not null"`   // 字典编码
		ItemName   string          `json:"item_name" gorm:"column:item_name;type:varchar(32);not null"` // 项名
		ItemCode   string          `json:"item_code" gorm:"column:item_code;type:varchar(32);not null"` // 项编码
		ItemValue  string          `json:"item_value" gorm:"column:item_value;type:varchar(255)"`       // 项值
		ItemValue1 string          `json:"item_value1" gorm:"column:item_value1;type:varchar(255)"`     // 项值1
		ItemValue2 string          `json:"item_value2" gorm:"column:item_value2;type:varchar(255)"`     // 项值2
		ItemValue3 string          `json:"item_value3" gorm:"column:item_value3;type:varchar(255)"`     // 项值3
		ItemValue4 string          `json:"item_value4" gorm:"column:item_value4;type:varchar(255)"`     // 项值4
		ItemDesc   string          `json:"item_desc" gorm:"column:item_desc;type:varchar(255)"`         // 项描述
		Order      int             `json:"order" gorm:"column:order;type:int"`                          // 排序
		CreatedAt  int64           `json:"created_at" gorm:"autoCreateTime:milli"`                      // 创建时间
		UpdatedAt  int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                      // 更新时间
		DeletedAt  *gorm.DeletedAt `json:"deleted_at"`                                                  // 删除时间
	}
)
