package etcdinteraction

import (
	"fmt"
)

func Fuzz(data []byte) int {
	l := len(data)
	fmt.Println(l)
	return 0
}
