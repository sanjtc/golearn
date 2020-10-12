package etcd

import (
	"log"
	"os"
	"path"
	"testing"
)

func TestEtcdInteraction(t *testing.T) {
	filePath, _ := os.Getwd()
	log.Println(filePath)
	dir := path.Dir(filePath)
	log.Println(dir)
	dir = path.Dir(dir)
	log.Println(dir)
}
