package common

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

//PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`           // 页码
	PageSize int `json:"page_size" form:"page_size"` // 每页大小
}

//GetById Find by id structure
type GetById struct {
	ID int64 `json:"id" form:"id"` // 主键ID
}

type IdsReq struct {
	Ids []int64 `json:"ids" form:"ids"`
}

type Empty struct{}
