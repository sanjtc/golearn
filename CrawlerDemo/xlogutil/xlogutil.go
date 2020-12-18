package xlogutil

import (
	"fmt"
	"log"
)

// func GetCodeLoc() (file string, line int) {
// 	_, file, line, _ = runtime.Caller(1)
// 	return
// }

func Error(v ...interface{}) {
	log.Println("E | ", fmt.Sprint(v...))
}

func Warning(v ...interface{}) {
	log.Println("W | ", fmt.Sprint(v...))
}
