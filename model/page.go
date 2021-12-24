package model

// Page info
type PageInfo struct {
	// Which page to display
	PageNum int `form:"page_num" binding:"omitempty,gt=0"`
	// How much data to display per page
	PageSize int `form:"page_size" binding:"omitempty,gt=0"`
	// The amount of data to skip
	SkipNum int
	// According to which field to sort
	SortField string `form:"sort_field" binding:"omitempty"`
	// Order of sort: ascending:asc descending:desc
	SortOrder string `form:"sort_order" binding:"omitempty"`
	// Is descending
	SortOrderIsDesc bool
}
