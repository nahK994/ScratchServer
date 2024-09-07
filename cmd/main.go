package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nahK994/SimpleServer/pkg/models"
	"github.com/nahK994/SimpleServer/pkg/server"
)

func main() {
	srv := server.Initiate("127.0.0.1:8000")
	srv.RequestHandler("/post", http.MethodPost, func(r models.Request, w *models.Response) {
		fmt.Println("TEST ===>", r.Body)
		w.StatusCode = http.StatusCreated
		w.Body = r.Body
	})

	srv.RequestHandler("/post", http.MethodGet, func(r models.Request, w *models.Response) {
		fmt.Println("TEST ===>", r.Body)
		w.StatusCode = http.StatusOK
		w.Body = "Hello World!!"
	})
	srv.RequestHandler("/get", http.MethodGet, func(r models.Request, w *models.Response) {
		fmt.Println("TEST ===>", r.Body)
		w.StatusCode = http.StatusOK
		w.Body = "Get request -> Hello World!!"
	})
	log.Fatal(srv.Start())
}
