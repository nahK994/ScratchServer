package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type node struct {
	name string
	net.Conn
}

var nodes []node

type Server struct {
	listenAddr string
	ln         net.Listener
	quitCh     chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitCh:     make(chan struct{}),
	}
}

func newNode(conn net.Conn) (*node, error) {
	conn.Write([]byte("Input name: "))
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	bufStr := strings.TrimSpace(string(buf[:n]))
	node := node{
		name: bufStr,
		Conn: conn,
	}
	return &node, nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		node, err := newNode(conn)
		if err != nil {
			conn.Write([]byte("cannot input user info"))
			conn.Close()
			return
		}
		nodes = append(nodes, *node)
		fmt.Printf("new connection to %s (%s)\n", node.name, conn.RemoteAddr())

		go s.readLoop(*node)
	}
}

func (s *Server) readLoop(node node) {
	buf := make([]byte, 2048)
	conn := node.Conn
	name := node.name
	msg := fmt.Sprintf("%s: ", node.name)
	conn.Write([]byte(msg))
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Close connection with %s\n", name)
			<-s.quitCh
		}

		fmt.Printf("%s: %s", name, string(buf[:n]))
		// receivedMsg := fmt.Sprintf("Received --> %s", string(buf[:n]))
		// conn.Write([]byte(receivedMsg))
		conn.Write([]byte(msg))
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

	return nil
}

func main() {
	fmt.Println("Server started")
	server := NewServer(":8000")
	log.Fatal(server.Start())
}
