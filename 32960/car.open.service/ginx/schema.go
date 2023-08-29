package ginx

// Validator .
type Validator interface {
	Validate() error
}

// StatusText 定义状态文本
type StatusText string

func (t StatusText) String() string {
	return string(t)
}

// 定义HTTP状态文本常量
const (
	OKStatus    StatusText = "OK"
	ErrorStatus StatusText = "ERROR"
	FailStatus  StatusText = "FAIL"
)

// StatusResult 状态结果
type StatusResult struct {
	Status StatusText `json:"status"`
}

// ErrorResult 响应错误
type ErrorResult struct {
	ErrorItem // 错误项
}

// ErrorItem 响应错误项
type ErrorItem struct {
	Status  StatusText `json:"staus"`
	Code    int        `json:"code"`    // 错误码
	Message string     `json:"message"` // 错误信息
}

// PaginationParam 分页查询条件
type PaginationParam struct {
	Pagination bool `form:"pagination,default=true"`                // 是否使用分页查询
	OnlyCount  bool `form:"-"`                                      // 是否仅查询count
	Current    uint `form:"current,default=1"`                      // 当前页
	PageSize   uint `form:"page_size,default=10" binding:"max=100"` // 页大小
}

// Pager .
type Pager interface {
	GetCurrent() uint
	GetPageSize() uint
	IsPagination() bool
}

// GetCurrent .
func (p *PaginationParam) GetCurrent() uint {
	return p.Current
}

// GetPageSize .
func (p *PaginationParam) GetPageSize() uint {
	return p.PageSize
}

// IsPagination .
func (p *PaginationParam) IsPagination() bool {
	return p.Pagination
}

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total    int  `json:"total"`
	Current  uint `json:"current"`
	PageSize uint `json:"page_size"`
}

// ListResult 响应列表数据
type ListResult struct {
	List       interface{}       `json:"list"`
	Pagination *PaginationResult `json:"pagination,omitempty"`
}

// SuccessResult 成功结果
type SuccessResult struct {
	Status StatusText  `json:"status"` // "OK"
	Data   interface{} `json:"data"`   // 返回数据
}
