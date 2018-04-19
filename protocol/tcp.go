package protocol

import (
	"dora.org/bara/gateway/util"
	"net"
	"time"
	"log"
	"runtime/debug"
	"io"
	"bytes"
)

type TCP struct {
	config TCPConfig
}

func (tcp *TCP) Run() (err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("bootstrap tcp proxy with err : %s \nstack: %s", err, string(debug.Stack()))
		}
	}()
	server := &util.Server{Address: *(tcp.config.Base.LocalAddress)}
	err = server.ListenTCP(tcp.tcpFunction)
	if err != nil {
		log.Println(err)
		log.Println(string(debug.Stack()))
	}
	return
}

func (tcp *TCP) Stop() error {
	return nil
}

func (tcp *TCP) tcpFunction(inConn *net.Conn) (e error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("tcp conn handler crashed with err : %s \nstack: %s", err, string(debug.Stack()))
		}
	}()
	outConn, e := net.DialTimeout("tcp", *tcp.config.Base.NextStage, time.Duration(*tcp.config.Base.Timeout)*time.Millisecond)
	log.Println("connect to next stage" + *tcp.config.Base.NextStage)
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
		var semaphore interface{}
		select {
		case semaphore = <-e1:
			//log.Printf("e1")
		case semaphore = <-e2:
			//log.Printf("e2")
		}
		dummy(semaphore)
		close(e1)
		close(e2)
	}()
	return
}

func dummy(dummy interface{}) {

}

func ioCopy(dst io.ReadWriter, src io.ReadWriter) (err error) {
	buf := make([]byte, 32*1024)
	var debug bytes.Buffer
	n := 0
	for {
		n, err = src.Read(buf)
		if n > 0 {
			debug.Write(buf[:n])
			log.Println(debug.String())
			if _, e := dst.Write(buf[0:n]); e != nil {
				return e
			}
		}
		if err != nil {
			return
		}
	}
}
