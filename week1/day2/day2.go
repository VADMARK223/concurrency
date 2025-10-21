package main

import (
	"fmt"
	"time"
)

func main() {
	const count = 3
	strChan := make(chan string, count)

	go func() {
		for i := 0; i < count; i++ {
			strChan <- fmt.Sprintf("task %d done", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(strChan)
	}()

	for msg := range strChan {
		fmt.Println(msg)
	}

	fmt.Println("All tasks done")
}
