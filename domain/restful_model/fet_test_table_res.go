package restful_model

type FetchTableRes struct {
	Records []TestTableRecord `json:"records"`
}

type TestTableRecord struct {
	Col1 string `json:"column1"`
	Col2 int    `json:"column2"`
}
