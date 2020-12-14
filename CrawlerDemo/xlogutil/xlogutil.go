package xlogutil

import (
	"log"
	"runtime"
)

func GetCodeLoc() (file string, line int) {
	_, file, line, _ = runtime.Caller(1)
	return
}

func Error(v interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("error: %v [file:%s, line:%d]", v, file, line)
}

func Warning(v interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("warning: %v [file:%s, line:%d]", v, file, line)
}

func Errorf(format string, v ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// log.Printf()
}
