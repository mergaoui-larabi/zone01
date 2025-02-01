package main

import (
	"fmt"
)

var RoutChannel = make(chan int)

func intger(ptr *int) {
	// _ = mtx
	for i := 0; i < 100000000; i++ {
		// mtx.Lock()
		*ptr++
		// mtx.Unlock()
	}
	RoutChannel <- *ptr
}

// func test(s string) {
// 	RoutChannel <- s
// }

func main() {
	// MainChannel := make(chan string)
	// Channel := make(chan string)
	// go func() {
	// 	MainChannel <- "main channel"
	// }()
	// go func() {
	// 	Channel <- "channel"
	// }()

	// select {
	// case msgone := <-MainChannel:
	// 	fmt.Println(msgone)
	// case msgsec := <-Channel:
	// 	fmt.Println(msgsec)
	// }
	// var waiter sync.WaitGroup
	// var mtx sync.Mutex

	i := 0

	go intger(&i)
	h := <-RoutChannel
	go intger(&h)
	g := <-RoutChannel

	fmt.Println(g)
}
