package main

import (
	"fmt"
	"time"
)

// the producer
func producer(name string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- "message from " + name
	}
}

// the consumer
func consumer(name string, c chan string) {
	for v := range c {
		fmt.Printf("%v got %v\n", name, v)
	}
}

func main() {
	var c = make(chan string)
	go producer("producer 1", c)
	go producer("producer 2", c)
	go producer("producer 3", c)
	go producer("producer 4", c)
	go producer("producer 5", c)

	go consumer("consumer 1", c)
	go consumer("consumer 2", c)
	go consumer("consumer 3", c)

	time.Sleep(1 * time.Second)
}
