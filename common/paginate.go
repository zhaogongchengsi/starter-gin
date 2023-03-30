package common

import (
	"gorm.io/gorm"
)

// Paginate 分页插件
//
// 使用方法看这:
//
// https://gorm.io/zh_CN/docs/scopes.html#%E5%88%86%E9%A1%B5
func Paginate(pageInfo Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := pageInfo.Page
		if page <= 0 {
			page = 1
		}

		pageSize := pageInfo.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
