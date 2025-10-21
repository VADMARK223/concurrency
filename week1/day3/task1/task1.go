package main

import (
	"concurrency/utils"
	"fmt"
	"time"
)

func worker(name string, ch chan string) {
	delay := time.Duration(utils.GetRandomInt(3)) * time.Second
	time.Sleep(delay)
	ch <- fmt.Sprintf("Worker %s done", name)
}

func main() {
	strChan := make(chan string)

	go worker("A", strChan)
	go worker("B", strChan)
	go worker("C", strChan)

	select {
	case msg := <-strChan:
		fmt.Println(msg, "first!")
	}

	fmt.Println("Done")
}
