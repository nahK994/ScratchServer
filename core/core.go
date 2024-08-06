package core

import (
	"fmt"
	"net"
)

type Server struct {
	ListenAddr string
	Ln         net.Listener
	QuitCh     chan struct{}
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

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Close connection with %s\n", conn.RemoteAddr())
			<-s.QuitCh
		}
		fmt.Printf("%s: %s", conn.RemoteAddr(), string(buf[:n]))
	}
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

// func main() {
// 	fmt.Println("Server started")
// 	server := NewServer(":8000")
// 	log.Fatal(server.Start())
// }
