package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)
	go updateMessage("test", &wg)
	wg.Wait()

	if msg != "test" {
		t.Errorf("Expected msg to be test")
	}
}

func Test_printMessage(t *testing.T) {

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	tempMsg := msg
	msg = "Hello, test!"

	printMessage()

	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	msg = tempMsg

	if !strings.Contains(output, "Hello, test!") {
		t.Errorf("Expected test to be printed")
	}
}
