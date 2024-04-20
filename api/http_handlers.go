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
		}

		req := restful_model.AddReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
		}

		svcAddRes, err := service.AddService(service_model.AddSvcReq{Num1: req.A, Num2: req.B})
		if err != nil {
			showError(writer, err, 500)
		}

		jsonRes, err := json.Marshal(restful_model.AddRes{Result: svcAddRes.Result})
		if err != nil {
			showError(writer, err, 500)
		}

		_, err = writer.Write(jsonRes)
		if err != nil {
			showError(writer, err, 500)
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
		}

		req := restful_model.SubReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
		}

		svcSubRes, err := service.SubService(service_model.SubSvcReq{Num1: req.A, Num2: req.B})
		if err != nil {
			showError(writer, err, 500)
		}

		jsonResSub, err := json.Marshal(restful_model.SubRes{Result: svcSubRes.Result})
		if err != nil {
			showError(writer, err, 500)

		}

		_, err = writer.Write(jsonResSub)
		if err != nil {
			showError(writer, err, 500)
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
		}

		req := restful_model.MulReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
		}

		svcMulRes, err := service.MulService(service_model.MulSvcReq{Num1: req.A, Num2: req.B})
		if err != nil {
			showError(writer, err, 500)

		}

		jsonResMul, err := json.Marshal(restful_model.MulRes{Result: svcMulRes.Result})
		if err != nil {
			showError(writer, err, 500)

		}

		_, err = writer.Write(jsonResMul)
		if err != nil {
			showError(writer, err, 500)
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
		}

		req := restful_model.DivReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
		}

		svcDivRes, err := service.DivService(service_model.DivSvcReq{Num1: req.A, Num2: req.B})
		if err != nil {
			showError(writer, err, 500)
		}

		jsonResDiv, err := json.Marshal(restful_model.DivRes{Result: svcDivRes.Result})
		if err != nil {
			showError(writer, err, 500)
		}

		_, err = writer.Write(jsonResDiv) //body response print
		if err != nil {
			showError(writer, err, 500)
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
		}

		req := restful_model.FetchTableReq{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			showError(writer, err, 400)
		}

		fetchTestTableSvcRes, err := service.FetchTestTableRecords(request.Context(), service_model.FetchTableSvcReq{})
		if err != nil {
			showError(writer, err, 500)
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
		}

		_, err = writer.Write(jsonResDiv)
		if err != nil {
			showError(writer, err, 500)
		}
	}
}
