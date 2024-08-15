package server

import (
	"log/slog"
	"net"

	"github.com/nahK994/ScratchServer/handlers"
	"github.com/nahK994/ScratchServer/utils"
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
	n, err := p.conn.Read(buf)
	if err != nil {
		slog.Error("peer read error", "err", err, "remoteAddr", p.conn.RemoteAddr())
		p.conn.Close()
		return
	}

	request := handlers.HandleRequest(buf[:n])
	//fmt.Println("Request from", p.conn.RemoteAddr(), " ==>", request)
	utils.RouteMapper[request.UrlPath](*request)
	handlers.HandleResponse(p.conn)
}
