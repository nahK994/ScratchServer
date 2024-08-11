package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/nahK994/ScratchServer/handlers"
)

type Peer struct {
	conn net.Conn
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
	}
}

func (p *Peer) readConn() {
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			slog.Error("peer read error", "err", err, "remoteAddr", p.conn.RemoteAddr())
			p.conn.Close()
			return
		}
		msg := fmt.Sprintf("server received -> %s", string(buf[:n]))
		p.conn.Write([]byte(msg))

		request := handlers.HandleRequest(buf[:n])
		fmt.Println("Final ==>", request)
	}
}
