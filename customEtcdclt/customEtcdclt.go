package main

import (
	"flag"
)

func main() {
	const httpMode = true

	if httpMode {
		StartHTTPListen(":8080")
	} else {
		flag.Parse()
		argv := flag.Args()
		ParseCommand(argv)
	}
}
