package protocol

import (
	"dora.org/bara/gateway/util"
	"strings"
	"strconv"
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
