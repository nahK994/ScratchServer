package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/server"
)

func main() {
	srv := server.Initiate("127.0.0.1:8000")
	srv.RequestHandler("/post", http.MethodPost, func(r models.Request, w *models.Response) {
		fmt.Println("TEST ===>", r.Body)
		w.StatusCode = 201
		w.Body = r.Body
	})

	srv.RequestHandler("/get", http.MethodGet, func(r models.Request, w *models.Response) {
		fmt.Println("TEST ===>", r.Body)
		w.StatusCode = 200
		w.Body = "Hello World!!"
	})
	log.Fatal(srv.Start())
}
