package api

import (
	"net/http"
)

type HTTPServer interface {
	Start() error
}

type myHttpServer struct {
	handlersMap map[string]func(http.ResponseWriter, *http.Request)
}

func NewHttpServer() HTTPServer {
	return myHttpServer{handlersMap: registerHandlers()} //Handlers container is filled by registerHandlers function
}

func (s myHttpServer) Start() error {
	mux := http.ServeMux{}
	for k, v := range s.handlersMap {
		mux.HandleFunc(k, v)
	}
	err := http.ListenAndServe(":8080", &mux) //start server//
	if err != nil {
		return err
	}
	return nil
}
