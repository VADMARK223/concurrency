package main

import (
	"fmt"
	"time"
)

func fetchData(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Data received!"
}

func main() {
	dataCh := make(chan string)
	go fetchData(dataCh)

	select {
	case msg := <-dataCh:
		fmt.Println(msg)
	case temp := <-time.After(2 * time.Second):
		fmt.Println("Timeout reached:", temp)
	}
}
