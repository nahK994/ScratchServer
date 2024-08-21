package main

import (
	"fmt"
	"log"

	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/server"
)

func main() {
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
