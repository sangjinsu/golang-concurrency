package main

import "fmt"

func main() {
	c := make(chan int)
	// send
	go send(c)

	// receive
	receive(c)

}

// send
func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// receive
func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
