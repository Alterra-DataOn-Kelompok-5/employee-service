package dto

import (
	"math"
)

// swagger:parameters FindEmployees
type Pagination struct {
	// in: query
	// items.page.pattern: \d+
	// items.page_size.pattern: \d+
	Page     *int `query:"page" json:"page"`
	PageSize *int `query:"page_size" json:"page_size"`
}

// swagger:parameters FindEmployees
type SearchGetRequest struct {
	// in: query
	// items.search.pattern: \w+
	Pagination
	Search   string   `query:"search"`
}

type SearchGetResponse[T any] struct {
	Data           []T `json:"data"`
	PaginationInfo PaginationInfo
}

type PaginationInfo struct {
	*Pagination
	Count       int  `json:"count"`
	MoreRecords bool `json:"more_records"`
	TotalPage   int  `json:"total_page"`
}

type ByIDRequest struct {
	ID uint `param:"id" validate:"required"`
}

func GetLimitOffset(p *Pagination) (limit, offset int) {

	if p.PageSize != nil {
		limit = *p.PageSize
	} else {
		limit = 10
		p.PageSize = &limit
	}

	if p.Page != nil {
		offset = (*p.Page - 1) * limit
	} else {
		offset = 0
	}

	return
}

func CheckInfoPagination(p *Pagination, count int64) *PaginationInfo {
	info := PaginationInfo{
		Pagination: p,
	}
	var page int
	if p.Page != nil {
		page = *p.Page
	} else {
		page = 1
	}
	info.Page = &page

	info.Count = int(count)
	info.TotalPage = int(math.Ceil(float64(count) / float64(*p.PageSize)))
	info.MoreRecords = true
	if *p.Page >= info.TotalPage {
		info.MoreRecords = false
	}

	return &info
}
