package main

import (
	"fmt"
	"time"
)

func main() {
	
	go work() //fork point
	
	fmt.Println("done waiting, main exists")
	
	//join point 
}

func work() {
	time.Sleep(100 * time.Microsecond)
	fmt.Println("task1 ")
}