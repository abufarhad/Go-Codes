package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})

	go func() {
		work() //fork point
		done <- struct{}{}
	}()

	<-done //join point
	fmt.Println(time.Since(now))
	fmt.Println("done waiting, main exists")
}

func work() {
	time.Sleep(500 * time.Microsecond)
	fmt.Println("task1 ")
	
}
