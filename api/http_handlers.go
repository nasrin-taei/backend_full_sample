package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"restful/domain/restful_model"
	"restful/domain/service_model"
	"restful/service"
)

var handlers map[string]func(http.ResponseWriter, *http.Request) = make(map[string]func(http.ResponseWriter, *http.Request))

const (
	getMethod    = "GET"
	postMethod   = "POST"
	putMethod    = "PUT"
	deleteMethod = "DELETE"
)

func registerHandlers() map[string]func(http.ResponseWriter, *http.Request) {
	add()
	sub()
	mul()
	div()
	fetchTestTable()
	addBook()
	return handlers
}

func showError(w http.ResponseWriter, err error, statusCode int) {
	errObj := restful_model.ErrorMessageResponse{Message: err.Error()}
	marErrObj, err := json.Marshal(errObj)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(statusCode)
	w.Write(marErrObj)
}

func showPanic(writer http.ResponseWriter) {
	if e := recover(); e != nil {
		m := restful_model.ErrorMessageResponse{Message: fmt.Sprint(e)}
		marshal, err := json.Marshal(m)
		if err != nil {
			return
		}
		writer.Write(marshal)
	}
}

func add() {
	handlers["/add"] = func(writer http.ResponseWriter, request *http.Request) {

		defer showPanic(writer)
		writer.Header().Add("Content-Type", "application/json")

		if request.Method != postMethod {
			showError(writer, errors.New("invalid request method"), 400)
			return
		}

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		req := restful_model.AddReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		svcAddRes, err := service.AddService(service_model.AddSvcReq{Num1: req.A, Num2: req.B})
		if err != nil {
			showError(writer, err, 500)
			return
		}

		jsonRes, err := json.Marshal(restful_model.AddRes{Result: svcAddRes.Result})
		if err != nil {
			showError(writer, err, 500)
			return
		}

		_, err = writer.Write(jsonRes)
		if err != nil {
			showError(writer, err, 500)
			return
		}
	}
}

func sub() {
	handlers["/sub"] = func(writer http.ResponseWriter, request *http.Request) {

		defer showPanic(writer)
		writer.Header().Add("Content-Type", "application/json")

		if request.Method != postMethod {
			showError(writer, errors.New("invalid request method"), 400)
			return
		}

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		req := restful_model.SubReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		svcSubRes, err := service.SubService(service_model.SubSvcReq{Num1: req.A, Num2: req.B})
		if err != nil {
			showError(writer, err, 500)
			return
		}

		jsonResSub, err := json.Marshal(restful_model.SubRes{Result: svcSubRes.Result})
		if err != nil {
			showError(writer, err, 500)
			return

		}

		_, err = writer.Write(jsonResSub)
		if err != nil {
			showError(writer, err, 500)
			return
		}
	}
}

func mul() {
	handlers["/mul"] = func(writer http.ResponseWriter, request *http.Request) {

		defer showPanic(writer)
		writer.Header().Add("Content-Type", "application/json")

		if request.Method != postMethod {
			showError(writer, errors.New("invalid request method"), 400)
			return
		}

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		req := restful_model.MulReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		svcMulRes, err := service.MulService(service_model.MulSvcReq{Num1: req.A, Num2: req.B})
		if err != nil {
			showError(writer, err, 500)
			return

		}

		jsonResMul, err := json.Marshal(restful_model.MulRes{Result: svcMulRes.Result})
		if err != nil {
			showError(writer, err, 500)
			return

		}

		_, err = writer.Write(jsonResMul)
		if err != nil {
			showError(writer, err, 500)
			return
		}

	}
}

func div() {
	handlers["/div"] = func(writer http.ResponseWriter, request *http.Request) {

		defer showPanic(writer)
		writer.Header().Add("Content-Type", "application/json")

		if request.Method != postMethod {
			showError(writer, errors.New("invalid request method"), 400)
			return
		}

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		req := restful_model.DivReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		svcDivRes, err := service.DivService(service_model.DivSvcReq{Num1: req.A, Num2: req.B})
		if err != nil {
			showError(writer, err, 500)
			return
		}

		jsonResDiv, err := json.Marshal(restful_model.DivRes{Result: svcDivRes.Result})
		if err != nil {
			showError(writer, err, 500)
			return
		}

		_, err = writer.Write(jsonResDiv) //body response print
		if err != nil {
			showError(writer, err, 500)
			return
		}

	}
}

func fetchTestTable() {
	handlers["/fetch_test_table"] = func(writer http.ResponseWriter, request *http.Request) {

		defer showPanic(writer)
		writer.Header().Add("Content-Type", "application/json")

		if request.Method != getMethod {
			showError(writer, errors.New("invalid request method"), 400)
			return
		}

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		req := restful_model.FetchTableReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		fetchTestTableSvcRes, err := service.FetchTestTableRecords(request.Context(), service_model.FetchTableSvcReq{})
		if err != nil {
			showError(writer, err, 500)
			return
		}

		resRecords := make([]restful_model.TestTableRecord, 0)
		for _, r := range fetchTestTableSvcRes.Records {
			t := restful_model.TestTableRecord{
				Col1: r.Col1,
				Col2: r.Col2,
			}
			resRecords = append(resRecords, t)
		}

		jsonResDiv, err := json.Marshal(restful_model.FetchTableRes{Records: resRecords})
		if err != nil {
			showError(writer, err, 500)
			return
		}

		_, err = writer.Write(jsonResDiv)
		if err != nil {
			showError(writer, err, 500)
			return
		}
	}
}

func addBook() {
	handlers["/add_book"] = func(writer http.ResponseWriter, request *http.Request) {

		defer showPanic(writer)
		writer.Header().Add("Content-Type", "application/json")

		if request.Method != postMethod {
			showError(writer, errors.New("invalid request method"), 400)
			return
		}

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		req := restful_model.AddBookReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
			return
		}

		_, err = service.AddBookService(request.Context(), service_model.AddBookSvcReq{
			Title:     req.Title,
			Count:     req.Count,
			UnitPrice: req.UnitPrice,
		})
		if err != nil {
			showError(writer, err, 400)
			return
		}

		marshal, err := json.Marshal(restful_model.AddBookRes{})
		if err != nil {
			showError(writer, err, 400)
			return
		}

		writer.Write(marshal)
		if err != nil {
			showError(writer, err, 500)
			return
		}
	}
}
