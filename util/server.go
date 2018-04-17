package util

import (
	"net"
	"log"
	"runtime/debug"
)

type Server struct {
	Ip   string
	Port int
}

func (server *Server) ListenTCP(function func(conn net.Conn) (error)) (err error) {
	tcpListener, e := net.ListenTCP("tcp", &net.TCPAddr{IP: []byte(server.Ip), Port: server.Port})
	if e != nil {
		err = e
		return
	} else {
		go func() {
			defer func() {
				if e := recover(); e != nil {
					log.Printf("ListenTCP crashed , err : %s , \ntrace:%s", e, string(debug.Stack()))
				}
			}()
			for {
				var conn net.Conn
				conn, err = tcpListener.AcceptTCP()
				if err == nil {
					go func() {
						function(conn)
					}()
				} else {
					log.Printf("TCP conn crashed , err : %s , \ntrace:%s", e, string(debug.Stack()))
					continue
				}
			}
		}()
	}
}
