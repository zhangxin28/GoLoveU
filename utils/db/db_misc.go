package db

import "database/sql"

// Paging 表示分页请求数据
type Paging struct {
	Page  int `json:"page"`  // 页码
	Limit int `json:"limit"` // 每页条数
	Total int `json:"total"` // 总数据条数
}

// Offset 表示取数据前要跳过多少数据
func (pp *Paging) Offset() int {
	offset := 0
	if pp.Page > 0 {
		offset = (pp.Page - 1) * pp.Limit
	}
	return offset
}

// TotalPage 表示总页数
func (pp *Paging) TotalPage() int {
	if pp.Total == 0 ||
		pp.Limit == 0 {
		return 0
	}
	totalPage := pp.Total / pp.Limit
	if pp.Total%pp.Limit > 0 {
		totalPage = totalPage + 1
	}
	return totalPage
}

// ParamPair 表示查询参数
type ParamPair struct {
	Query string        // 查询
	Args  []interface{} // 参数
}

// OrderByCol 表示排序信息
type OrderByCol struct {
	Column string // 排序字段
	Asc    bool   // 是否正序
}

// PageResult 表示分页返回数据
type PageResult struct {
	Page    *Paging     `json:"page"`    // 分页信息
	Results interface{} `json:"results"` // 数据
}

// CursorResult 表示Cursor分页返回数据
type CursorResult struct {
	Results interface{} `json:"results"` // 数据
	Cursor  string      `json:"cursor"`  // 下一页
}

// SQLNullString 表示一个字符串值可能为null
func SQLNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  len(value) > 0,
	}
}
