package restful_model

type AddBookReq struct {
	Title     string `json:"title"`
	Count     int    `json:"count"`
	UnitPrice int64  `json:"unitPrice"`
}
