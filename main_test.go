package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain_Execute(t *testing.T) {
	bkpStdOut := os.Stdout
	read, write, _ := os.Pipe()
	os.Stdout = write

	main()

	write.Close()
	out, _ := ioutil.ReadAll(read)
	os.Stdout = bkpStdOut

	if !strings.Contains(string(out), "FINISHED") {
		t.Error("Main not running")
	}
}
