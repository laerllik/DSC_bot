package logger

import (
	"log"
)

type logger struct {
	prefix string
}

func (l *logger) log(message interface{}) {
	log.Println(l.prefix, message)
}

func (l *logger) configure(prefix string) {
	l.prefix = prefix
}

var l = logger{prefix: ""}

func Log(message interface{}) {
	l.log(message)
}

func Configure(prefix string) {
	l.configure(prefix)
}
