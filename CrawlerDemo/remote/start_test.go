package main

import "testing"

func TestStart(t *testing.T) {
	UploadSrc()
}

func TestInterrupt(t *testing.T) {
	if err := ProcessInterrupt(); err != nil {
		t.Fatal(err)
	}
}
