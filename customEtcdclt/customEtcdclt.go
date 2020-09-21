package main

import (
	"flag"
)

func main() {
	const httpMode = true

	if httpMode {
		HTTPClient(":8080")
	} else {
		flag.Parse()
		argv := flag.Args()
		ParseCommand(argv)
	}
}
