package main

import (
	"fmt"
	"log"

	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/server"
)

func main() {
	srv := server.Initiate("127.0.0.1:8000")
	srv.RequestHandler("/login", "POST", func(r models.Request, w *models.Response) {
		fmt.Println("TEST ===>", r.Body)
		w.StatusCode = 201
		w.Body = r.Body
	})
	log.Fatal(srv.Start())
}
