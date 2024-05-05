package service_model

type UpdateBookSvcReq struct {
	Title     string
	Count     int
	UnitPrice int64
	Id        int64
}
