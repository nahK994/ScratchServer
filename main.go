package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("new connection to", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Close connection with %s\n", conn.RemoteAddr())
			break
		}

		msg := buf[:n]
		fmt.Printf("%s: %s\n", conn.RemoteAddr(), string(msg))
		conn.Write([]byte(fmt.Sprintf("Received --> %s\n", string(msg))))
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	s.ln = ln
	s.acceptLoop()
	<-s.quitch

	return nil
}

func main() {
	server := NewServer(":8000")
	log.Fatal(server.Start())
}
