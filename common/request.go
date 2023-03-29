package common

type Page struct {
	Page     int `query:"page" json:"page" form:"page"`             // 页码
	PageSize int `query:"pageSize" json:"pageSize" form:"pageSize"` // 每页大小
}
