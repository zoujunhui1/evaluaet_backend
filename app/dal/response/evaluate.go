package response

type GetProductListResp struct {
	Total    int64      `json:"total"`
	PageSize int        `json:"page_size"`
	Page     int        `json:"page"`
	List     []*Product `json:"list"`
}

type Product struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
