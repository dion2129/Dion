package main

import "fmt"

func sendData(ch chan string) {
	ch <- "Hello"
	ch <- "ffrom"
	ch <- "goroutine"

	close(ch)
}

func main() {
	ch := make(chan string, 3)

	go sendData(ch)

	for value := range ch {
		fmt.Println(value)
	}
}