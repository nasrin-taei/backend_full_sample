package restful_model

type UpdateBookReq struct {
	Title     string `json:"title"`
	Count     int    `json:"count"`
	UnitPrice int64  `json:"unitPrice"`
	Id        int64  `json:"id"`
}
