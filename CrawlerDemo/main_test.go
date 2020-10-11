package main

import (
	"log"
	"os"
	"testing"
)

func TestFunc(t *testing.T) {
	path, _ := os.Getwd()
	log.Println(path)
}
