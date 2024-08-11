package server

import (
	"fmt"
	"log/slog"
	"net"
)

type Peer struct {
	conn net.Conn
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
	}
}

func (peer *Peer) handleConn() {

	slog.Info("new peer connected", "remoteAddr", peer.conn.RemoteAddr())

	if err := peer.readConn(); err != nil {
		slog.Error("peer read error", "err", err, "remoteAddr", peer.conn.RemoteAddr())
		peer.conn.Close()
	}
}

func (p *Peer) readConn() error {
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			return err
		}
		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])
		msg := fmt.Sprintf("server received -> %s", string(msgBuf))
		p.conn.Write([]byte(msg))
	}
}
