package utils

import (
	"net/url"
	"strconv"
)

// Pagination represents how to paginate over results
type Pagination struct {
	Max    int `json:"max"`
	Offset int `json:"offset"`
}

// PagedResult represents a partial result
type PagedResult struct {
	Data  []interface{} `json:"data"`
	Total int16         `json:"total"`
}

// GetPagination extracts pagination params from URL query params
func GetPagination(vars url.Values) Pagination {
	max, _ := strconv.Atoi(vars.Get("max"))
	off, _ := strconv.Atoi(vars.Get("offset"))

	return Pagination{Max: max, Offset: off}
}
