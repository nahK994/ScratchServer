package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nahK994/ScratchServer/http/handlers"
)

type Server struct {
	ListenAddr string
	Ln         net.Listener
	QuitCh     chan struct{}
}

func handleGetRequest(conn net.Conn) {
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Connection: close\r\n\r\n" +
		"Hello, World!"
	conn.Write([]byte(response))
	conn.Close()
}

func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr: listenAddr,
		QuitCh:     make(chan struct{}),
	}
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.Ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Printf("new connection to %s\n", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	buf := make([]byte, 2048)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Close connection with %s\n", conn.RemoteAddr())
		conn.Close()
		return
	}

	// fmt.Printf("%s: %s", conn.RemoteAddr(), string(buf[:n]))
	handlers.HandleRequest(string(buf[:n]))
	handleGetRequest(conn)
	fmt.Printf("connection closed with %s\n", conn.RemoteAddr())
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	s.Ln = ln
	s.acceptLoop()

	return nil
}

func main() {
	fmt.Println("Server started")
	server := NewServer(":8000")
	log.Fatal(server.Start())
}
