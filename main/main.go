package main

import (
	"restful/api"
	"restful/repository"
)

func main() {
	err := repository.ConnectToPostgres()
	if err != nil {
		return
	}

	err = api.NewHttpServer().Start()
	if err != nil {
		return
	}
}
