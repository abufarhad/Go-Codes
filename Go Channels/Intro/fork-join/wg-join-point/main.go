package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		work() //fork point
	}()

	wg.Wait() //join point
	fmt.Println(time.Since(now))
	fmt.Println("done waiting, main exists")
}

func work() {
	time.Sleep(500 * time.Microsecond)
	fmt.Println("task1 ")
}
