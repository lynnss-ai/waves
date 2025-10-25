// Package dbcore 数据库核心模型
//
// @File: basemodel.go
// @Desc: 数据库核心模型
// @Date: 2025-10-25

package dbcore

import (
	"context"

	"gorm.io/gorm"
)

type BaseModel[T any] struct {
	DB *gorm.DB
}

// 事务绑定
func (m *BaseModel[T]) WithTrans(ctx context.Context) *BaseModel[T] {
	return &BaseModel[T]{DB: GetDB(ctx, m.DB)}
}

// 插入数据
func (m *BaseModel[T]) Insert(ctx context.Context, v *T) error {
	return m.DB.WithContext(ctx).Create(v).Error
}

// 批量插入数据
func (m *BaseModel[T]) InsertBatch(ctx context.Context, v []*T, batchSize int) error {
	return m.DB.WithContext(ctx).CreateInBatches(v, batchSize).Error
}

// 根据主键ID更新数据
func (m *BaseModel[T]) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.DB.WithContext(ctx).Model(new(T)).Where("id = ?", id).Updates(v).Error
}

// 根据主键ID删除数据
func (m *BaseModel[T]) Delete(ctx context.Context, id string) error {
	return m.DB.WithContext(ctx).Delete(new(T), id).Error
}

// 根据主键ID查询数据
func (m *BaseModel[T]) Find(ctx context.Context, id string) (*T, error) {
	var v T
	err := m.DB.WithContext(ctx).Where("id = ?", id).First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// 根据主键ID列表查询数据
func (m *BaseModel[T]) ListByIds(ctx context.Context, ids []string) ([]*T, error) {
	var list []*T
	err := m.DB.WithContext(ctx).Where("id IN ?", ids).Find(&list).Error
	return list, err
}

// 根据条件查询数据
// 特别注意: 这种方式只支持 “字段 = 值” 的等值匹配
func (m *BaseModel[T]) FindBy(ctx context.Context, cond map[string]interface{}) (*T, error) {
	var v T
	err := m.DB.WithContext(ctx).Where(cond).First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// 根据条件查询数据是否存在
func (m *BaseModel[T]) Exist(ctx context.Context, cond map[string]interface{}) (bool, error) {
	var count int64
	err := m.DB.WithContext(ctx).Model(new(T)).Where(cond).Count(&count).Error
	return count > 0, err
}

type Cond struct {
	Query string
	Args  []interface{}
	Map   map[string]interface{}
}

func (m *BaseModel[T]) FindByCond(ctx context.Context, cond Cond) (*T, error) {
	var v T
	db := m.DB.WithContext(ctx)

	if len(cond.Map) > 0 {
		db = db.Where(cond.Map)
	}
	if cond.Query != "" {
		db = db.Where(cond.Query, cond.Args...)
	}

	err := db.First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// 查询数据总数
func (m *BaseModel[T]) count(ctx context.Context) (int64, error) {
	var count int64
	err := m.DB.WithContext(ctx).Model(new(T)).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
