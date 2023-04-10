package common

// Page 翻页器
type Page struct {
	Page     int `query:"page" json:"page" form:"page"`             // 页码
	PageSize int `query:"pageSize" json:"pageSize" form:"pageSize"` // 每页大小
}

// IDSet 获取id集合
type IDSet[A int | uint | uint8 | string | float64 | float32] struct {
	Ids []A `json:"ids" from:"ids" query:"pageSize"`
}

type ID[I int | uint | uint8 | uint16 | float32 | float64] struct {
	Id I `uri:"id" json:"id" query:"id"`
}
