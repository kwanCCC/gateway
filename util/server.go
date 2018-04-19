package util

import (
	"net"
	"log"
	"runtime/debug"
)

type Server struct {
	Address string
}

func (server *Server) ListenTCP(function func(conn net.Conn) (error)) (err error) {
	var tcpAddr *net.TCPAddr
	tcpAddr, err = net.ResolveTCPAddr("tcp", server.Address)
	tcpListener, e := net.ListenTCP("tcp", tcpAddr)
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
				conn, err := tcpListener.AcceptTCP()
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
	return
}
