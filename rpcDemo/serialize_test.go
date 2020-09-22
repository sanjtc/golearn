package main

import (
	"encoding/json"
	"log"
	"os"
	"testing"
	"unsafe"

	"github.com/pantskun/golearn/rpcDemo/proto/example"
)

func TestSerialize(t *testing.T) {
	test := &example.Example{
		StringVal: "stringval",
		BytesVal:  []byte{0, 0, 1},
		EmbeddedExample: &example.Example_EmbeddedMessage{
			Int32Val:  1,
			StringVal: "embeddedStringval",
		},
		RepeatedInt32Val:  []int32{2},
		RepeatedStringVal: []string{"repeatedStringValue"},
	}

	testBytes, _ := json.MarshalIndent(test, "", "\t")

	file, err := os.Create("test.json")
	if err != nil {
		log.Println(err)
		return
	}

	info, _ := file.Stat()

	_, _ = file.WriteAt(testBytes, info.Size())
}

func TestOneof(t *testing.T) {
	test := &example.Example{}
	test.TestOneof = &example.Example_OneofEmbeddedExample{
		&example.Example_EmbeddedMessage{Int32Val: 1, StringVal: "OneofEmbeddedExample"}}
	t.Log(test.TestOneof)
	test.TestOneof = &example.Example_OneofName{OneofName: "OneofName"}
	t.Log(test.TestOneof)
}

type testStruct struct {
	intval    int
	stringval string
}

func TestUnsafe(t *testing.T) {
	test := testStruct{1, "2"}
	testint := 2
	t.Log(unsafe.Sizeof(test))
	t.Log(unsafe.Alignof(testint))
}
