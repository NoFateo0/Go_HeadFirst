package main

import (
	"fmt"
	"time"
)

func a(channel chan bool) {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
	channel <- true
}

func b(channel chan bool) {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
	channel <- true
}

func main() {
	gg := make(chan bool)
	go a(gg)
	go b(gg)
	time.Sleep(time.Second)
	<-gg
	<-gg
	fmt.Println("end main()")
}
