package protocol

import (
	"runtime/debug"
	"fmt"
)

type Protocol interface {
	Run() error
	Stop() error
}

type Bootstrap struct {
	Name     string
	Protocol Protocol
}

var ready = make(map[string]*Bootstrap)

func Register(name string, protocol Protocol) (status bool) {
	if _, ok := ready[name]; ok {
		status = false
		fmt.Errorf("protocol %s has been registered already\n", name)
	} else {
		ready[name] = &Bootstrap{name, protocol}
		fmt.Printf("protocol %s register successfully\n", name)
		status = true
	}
	return
}

func Boot(name string) (err error) {
	if bootstrap, ok := ready[name]; ok {
		go func() {
			e := bootstrap.Protocol.Run()
			if e != nil {
				err = e
				fmt.Errorf("setup %s fail and trace %s \n", name, string(debug.Stack()))
			} else {
				fmt.Printf("setup %s successfully \n", name)
			}
		}()
		return
	}
	return
}
