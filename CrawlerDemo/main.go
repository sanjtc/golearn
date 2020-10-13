package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/pantskun/golearn/CrawlerDemo/pathutils"
)

func main() {
	mainPath := pathutils.GetModulePath() + "/main/main.go"
	cmd1 := exec.Command("go", "run", mainPath)

	var out bytes.Buffer
	cmd1.Stdout = &out

	err := cmd1.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}
