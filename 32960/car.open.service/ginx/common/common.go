package common

// PaginationParam 分页查询条件
type PaginationParam struct {
	Pagination bool `json:"pagination" form:"pagination"`                            // 是否使用分页查询
	OnlyCount  bool `json:"-" form:"-"`                                              // 是否仅查询count
	Current    int  `json:"current" form:"current,default=1"`                        // 当前页
	PageSize   int  `json:"page_size" form:"page_size,default=10" binding:"max=100"` // 页大小
}

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total    int `json:"total_records"` // 条数
	Current  int `json:"current_page"`  // 当前页
	PageSize int `json:"page_size"`     // 页大小
}

// Page 分页
func (o *PaginationParam) Page() {
	if o.PageSize > 50 || o.PageSize == 0 {
		o.PageSize = 50
	}
}
