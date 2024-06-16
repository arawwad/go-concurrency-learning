package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, waitGroup *sync.WaitGroup) {
	msg = s
	defer waitGroup.Done()
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	var waitGroup sync.WaitGroup

	msg = "Hello, world!"

	waitGroup.Add(1)
	go updateMessage("Hello, universe!", &waitGroup)
	waitGroup.Wait()
	printMessage()

	waitGroup.Add(1)
	go updateMessage("Hello, cosmos!", &waitGroup)
	waitGroup.Wait()
	printMessage()

	waitGroup.Add(1)
	go updateMessage("Hello, world!", &waitGroup)
	waitGroup.Wait()
	printMessage()
}
