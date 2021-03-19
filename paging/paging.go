package paging

import (
	"math"
)

// Constants
const (
	DefaultLimit int64 = 25
)

// Paging struct
type Paging struct {
	CurrentPage int64 `json:"current_page"`
	Total       int64 `json:"total"`
	TotalPage   int64 `json:"total_page"`
	Limit       int64 `json:"limit"`
	Skip        int64 `json:"skip"`
}

// New paging object
func New(page int64, pageSize int64, total int64) *Paging {
	var pageInfo Paging
	limit := DefaultLimit

	if pageSize > 0 && pageSize <= limit {
		pageInfo.Limit = pageSize
	} else {
		pageInfo.Limit = limit
	}

	totalPage := int64(math.Ceil(float64(total) / float64(pageInfo.Limit)))
	pageInfo.Total = total
	pageInfo.TotalPage = totalPage
	if page < 1 || totalPage == 0 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}

	pageInfo.CurrentPage = page
	pageInfo.Skip = (page - 1) * pageInfo.Limit
	return &pageInfo
}
