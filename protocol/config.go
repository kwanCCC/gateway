package protocol

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

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

func InitConfig() {
	tcpConfig = &TCPConfig{}
	// flag parse
	app := kingpin.New("gateway", "Let's be gopher")
	app.Author(`666chan`)
	tcpCommand := app.Command("tcp", "tcp config")
	localAddress := tcpCommand.Flag("local", "listen local address").Required().Short('l').String()
	nextStage := tcpCommand.Flag("next", "proxy target").Required().Short('n').String()
	timeout := tcpCommand.Flag("timeout", "dial time out").Short('t').Default("5").Int()
	//parse args
	protocol := kingpin.MustParse(app.Parse(os.Args[1:]))
	base := &base{nextStage, localAddress, timeout}
	tcpConfig.Base = base
	if tcpConfig != nil {
		Register("tcp", &TCP{*tcpConfig})
	}
	Boot(protocol)

}
