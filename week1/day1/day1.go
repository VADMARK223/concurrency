package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go downloadFile(fmt.Sprintf("(File %d)", i), &wg)
	}

	wg.Wait()
	fmt.Println("Done")
}

func downloadFile(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Downloading  start...", name)
	time.Sleep(time.Second)

	fmt.Println("Completed", name)
}
