package db

import (
	"github.com/jinzhu/gorm"
	"goloveu/utils"
)

// SQLCnd represents the sql selector
type SQLCnd struct {
	SelectCols []string     //要查询的字段，如果为空，表示查询所有的字段
	Params     []ParamPair  //参数
	Orders     []OrderByCol //排序
	Paging     *Paging      //分页
}

// NewSQLCnd represents columns to query
func NewSQLCnd(selectCols ...string) *SQLCnd {
	s := &SQLCnd{}
	if len(selectCols) > 0 {
		s.SelectCols = append(s.SelectCols, selectCols...)
	}
	return s
}

// Where represents the where clause
func (s *SQLCnd) Where(query string, args ...interface{}) *SQLCnd {
	s.Params = append(s.Params, ParamPair{Query: query, Args: args})
	return s
}

// Eq represents equal where clause
func (s *SQLCnd) Eq(column string, args ...interface{}) *SQLCnd {
	s.Where(column+" = ?", args)
	return s
}

// NotEq represents not equal where clause
func (s *SQLCnd) NotEq(column string, args ...interface{}) *SQLCnd {
	s.Where(column+" <> ?", args)
	return s
}

// Gt represents grater than where clause
func (s *SQLCnd) Gt(column string, args ...interface{}) *SQLCnd {
	s.Where(column+" > ?", args)
	return s
}

// Gte represents grater and equal than where clause
func (s *SQLCnd) Gte(column string, args ...interface{}) *SQLCnd {
	s.Where(column+" >= ?", args)
	return s
}

// Lt represents less than where clause
func (s *SQLCnd) Lt(column string, args ...interface{}) *SQLCnd {
	s.Where(column+" < ?", args)
	return s
}

// Lte represents less equal than where clause
func (s *SQLCnd) Lte(column string, args ...interface{}) *SQLCnd {
	s.Where(column+" <= ?", args)
	return s
}

// Like represents like where clause: like contains
func (s *SQLCnd) Like(column string, str string) *SQLCnd {
	s.Where(column+" Like ?", "%"+str+"%")
	return s
}

// LikeStart represents like where clause: like start with
func (s *SQLCnd) LikeStart(column string, str string) *SQLCnd {
	s.Where(column+" Like ?", str+"%")
	return s
}

// LikeEnd represents like where clause: like end with
func (s *SQLCnd) LikeEnd(column string, str string) *SQLCnd {
	s.Where(column+" Like ?", "%"+str)
	return s
}

// In represents in where clause
func (s *SQLCnd) In(column string, args ...interface{}) *SQLCnd {
	s.Where(column+" in (?)", args)
	return s
}

// Asc represents column direction: asc
func (s *SQLCnd) Asc(column string) *SQLCnd {
	s.Orders = append(s.Orders, OrderByCol{Column: column, Asc: true})
	return s
}

// Desc represents column direction: desc
func (s *SQLCnd) Desc(column string) *SQLCnd {
	s.Orders = append(s.Orders, OrderByCol{Column: column, Asc: false})
	return s
}

// Limit represents the first page paging
func (s *SQLCnd) Limit(limit int) *SQLCnd {
	s.Page(1, limit)
	return s
}

// Page represents the paging parameters
func (s *SQLCnd) Page(page, limit int) *SQLCnd {
	if s.Paging == nil {
		s.Paging = &Paging{Page: page, Limit: limit}
	} else {
		s.Paging.Page = page
		s.Paging.Limit = limit
	}
	return s
}

// Build represents a whole sql
func (s *SQLCnd) Build(db *gorm.DB) *gorm.DB {
	ret := db

	// selector
	if len(s.SelectCols) > 0 {
		ret = ret.Select(s.SelectCols)
	}

	// where
	if len(s.Params) > 0 {
		for _, param := range s.Params {
			ret = ret.Where(param.Query, param.Args...)
		}
	}

	// order
	if len(s.Orders) > 0 {
		for _, order := range s.Orders {
			if order.Asc {
				ret = ret.Order(order.Column + " ASC")
			} else {
				ret = ret.Order(order.Column + " DESC")
			}
		}
	}

	// limit
	if s.Paging != nil && s.Paging.Limit > 0 {
		ret = ret.Limit(s.Paging.Limit)
	}

	// offset
	if s.Paging != nil && s.Paging.Offset() > 0 {
		ret = ret.Offset(s.Paging.Offset())
	}

	return ret
}

// Find represents to find records that match given conditions
func (s *SQLCnd) Find(db *gorm.DB, out interface{}) {
	if err := s.Build(db).Find(out).Error; err != nil {
		utils.LogError(err)
	}
}

// FindOne represents to find record that match given conditions
func (s *SQLCnd) FindOne(db *gorm.DB, out interface{}) {
	if err := s.Limit(1).Build(db).Find(out).Error; err != nil {
		utils.LogError(err)
	}
}

// Count represents to count that match given conditions
func (s *SQLCnd) Count(db *gorm.DB, model interface{}) int {
	ret := db.Model(model)

	// where
	if len(s.Params) > 0 {
		for _, query := range s.Params {
			ret = ret.Where(query.Query, query.Args...)
		}
	}

	var count int
	if err := ret.Count(&count).Error; err != nil {
		utils.LogError(err)
	}
	return count
}
