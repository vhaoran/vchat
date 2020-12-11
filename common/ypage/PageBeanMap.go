package ypage

import (
	"go.mongodb.org/mongo-driver/bson"
)

type PageBeanMap struct {
	PageNo      int64 `json:"page_no,omitempty"   bson:"page_no,omitempty"`
	RowsPerPage int64 `json:"rows_per_page,omitempty"   bson:"rows_per_page,omitempty"`
	PagesCount  int64 `json:"pages_count,omitempty"`
	RowsCount   int64 `json:"rows_count,omitempty"   bson:"rows_count,omitempty"`

	Where bson.M      `json:"where,omitempty"   bson:"where,omitempty"`
	Sort  bson.M      `json:"sort,omitempty"   bson:"sort,omitempty"`
	Data  interface{} `json:"data,omitempty"   bson:"data,omitempty"`
}

func (r *PageBeanMap) ValidateAdjust() {
	if r.PageNo < 1 {
		r.PageNo = 1
	}
	if r.RowsPerPage <= 1 {
		r.RowsPerPage = 10
	}
}

func (r *PageBeanMap) GetSkip() int64 {
	return (r.PageNo - 1) * r.RowsPerPage
}

func (r *PageBeanMap) GetPagesCount() int64 {
	all, rows := r.RowsCount, r.RowsPerPage
	if rows <= 0 {
		return 0
	}

	i := all % rows
	j := all / rows
	if i > 0 {
		return j + 1
	}
	return j
}
