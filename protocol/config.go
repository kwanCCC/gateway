package protocol

import "gopkg.in/alecthomas/kingpin.v2"

type base struct {
	NextStage    *string
	LocalAddress *string
	Timeout      *int
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
	nextStage := app.Flag("next", "proxy target").Required().Short('n').String()
	timeout := app.Flag("timeout", "dial time out").Short('t').Default("5").Int()
	base := &base{nextStage, localAddress, timeout}
	tcpConfig.Base = base
}

func Setup() {
	if tcpConfig != nil {
		Register("TCP", &TCP{*tcpConfig})
	}
}
