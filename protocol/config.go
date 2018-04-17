package protocol

import "gopkg.in/alecthomas/kingpin.v2"

type base struct {
	LocalAddress  *string
}

type TCPConfig struct {
	Base *base
}

type HTTPConfig struct {
	Base *base
}

var tcpConfig *TCPConfig

func init() {
	tcpConfig = &TCPConfig{}
	app := kingpin.New("gateway", "Let's be gopher")
	localAddress := app.Flag("local", "listen local address").Required().Short('l').String()
	base := &base{localAddress}
	tcpConfig.Base = base
}

func Setup() {
	if tcpConfig != nil {
		Register("TCP", &TCP{*tcpConfig})
	}
}
