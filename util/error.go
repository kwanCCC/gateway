package util

import (
	"bytes"
)

type Error struct {
	Info  string
	Stack string
}

var ERROR = []byte("error")
var STACK = []byte("stack")
var ENTER = []byte("\n")

func (err *Error) Error() string {
	var buf bytes.Buffer
	buf.Write(ERROR)
	buf.Write([]byte(err.Info))
	buf.Write(ENTER)
	buf.Write(STACK)
	buf.Write([]byte(err.Stack))
	buf.Write(ENTER)
	return buf.String()
}
