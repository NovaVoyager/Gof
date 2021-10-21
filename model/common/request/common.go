package request

//GetById Find by id structure
type GetById struct {
	ID int64 `json:"id" form:"id"` // 主键ID
}

type IdsReq struct {
	Ids []int64 `json:"ids" form:"ids"`
}
