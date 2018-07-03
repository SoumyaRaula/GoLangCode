package main

import (
	"fmt"
	"math/rand"
)

func new(done chan bool) {
	fmt.Println("A number from 1-100", rand.Intn(100))
	done <- true
}

func hello(newDone chan bool, done chan bool) {
	<-newDone
	fmt.Println("Hello world goroutine")
	done <- true
}
func main() {
	newDone := make(chan bool)
	done := make(chan bool)
	go hello(newDone, done)
	go new(newDone)
	fmt.Println("main function 1")
	<-done
	fmt.Println("main function 2")
}
