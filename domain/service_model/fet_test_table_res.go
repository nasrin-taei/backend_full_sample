package service_model

type FetchTableSvcRes struct {
	Records []TestTableSvcRecord
}

type TestTableSvcRecord struct {
	Col1 string
	Col2 int
}
