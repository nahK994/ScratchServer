package main

import (
	"fmt"
	"log"

	"github.com/nahK994/ScratchServer/models"
	"github.com/nahK994/ScratchServer/server"
)

func main() {
	srv := server.NewServer(server.Config{
		ListenAddress: "127.0.0.1:8000",
	})
	server.RegisterHandleFunc("/login", "POST", func(r models.Request, w *models.Response) {
		fmt.Println("TEST ===>", r.Body)
		w.StatusCode = 201
		w.Body = r.Body
	})
	log.Fatal(srv.Start())
}
