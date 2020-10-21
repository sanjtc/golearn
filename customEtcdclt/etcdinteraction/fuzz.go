package etcdinteraction

func Fuzz(data []byte) int {
	l := len(data)
	if l > 100 {
		panic(l)
	}
	return 0
}
