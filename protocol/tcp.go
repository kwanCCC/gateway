package protocol

import (
	"dora.org/bara/gateway/util"
	"strings"
	"strconv"
	"net"
	"time"
	"log"
	"runtime/debug"
	"io"
)

type TCP struct {
	config TCPConfig
}

func (tcp *TCP) Run() (err error) {
	hostAndIp := strings.Split(*(tcp.config.Base.LocalAddress), ":")
	if len(hostAndIp) != 2 {
		err = &util.Error{Info: "local address illegal"}
		return
	}
	port, e := strconv.Atoi(hostAndIp[1])
	if e != nil {
		err = &util.Error{Info: "local address illegal"}
		return
	}
	server := &util.Server{Ip: hostAndIp[0], Port: port}
	server.ListenTCP()
	return
}

func (tcp *TCP) Stop() error {
	return nil
}

func (tcp *TCP) tcpFunction(inConn *net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("tcp conn handler crashed with err : %s \nstack: %s", err, string(debug.Stack()))
		}
	}()
	outConn, e := net.DialTimeout("tcp", *tcp.config.Base.NextStage, time.Duration(*tcp.config.Base.Timeout)*time.Millisecond)
	if e != nil {
		return
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("bind crashed %s", err)
			}
			(*inConn).Close()
			outConn.Close()
		}()
		e1 := make(chan interface{}, 1)
		e2 := make(chan interface{}, 1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("bind crashed %s", err)
				}
			}()
			//_, err := io.Copy(dst, src)
			err := ioCopy(outConn, *inConn)
			e1 <- err
		}()
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("bind crashed %s", err)
				}
			}()
			//_, err := io.Copy(src, dst)
			err := ioCopy(*inConn, outConn)
			e2 <- err
		}()
		var err interface{}
		select {
		case err = <-e1:
			//log.Printf("e1")
		case err = <-e2:
			//log.Printf("e2")
		}
	}()
}

func ioCopy(dst io.ReadWriter, src io.ReadWriter) (err error) {
	buf := make([]byte, 32*1024)
	n := 0
	for {
		n, err = src.Read(buf)
		if n > 0 {
			if _, e := dst.Write(buf[0:n]); e != nil {
				return e
			}
		}
		if err != nil {
			return
		}
	}
}
