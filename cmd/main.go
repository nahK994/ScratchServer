package main

import (
	"fmt"
	"log"

	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/server"
)

func example_resp() {
	srv := server.InitiateResp("127.0.0.1:8000")
	srv.RequestHandler(func(request models.RespRequest, response *models.RespResponse) {
		response.Response = "+OK\r\n"
	})
	log.Fatal(srv.Start())
}

func example_http() {
	srv := server.InitiateHttp("127.0.0.1:8000")
	srv.RequestHandler("/post", "POST", func(r models.HttpRequest, w *models.HttpResponse) {
		fmt.Println("TEST ===>", r.Body)
		w.StatusCode = 201
		w.Body = r.Body
	})

	srv.RequestHandler("/get", "GET", func(r models.HttpRequest, w *models.HttpResponse) {
		fmt.Println("TEST ===>", r.Body)
		w.StatusCode = 200
		w.Body = "Hello World!!"
	})
	log.Fatal(srv.Start())
}

func main() {
	// example_http()
	example_resp()
}
