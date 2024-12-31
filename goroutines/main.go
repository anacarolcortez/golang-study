package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan string) {
	done <- "Hello!" + phrase // a value is sent through the channel when the function has completed
}

func slowGreet(phrase string, done chan string) {
	time.Sleep(3 * time.Second)
	done <- "Hello!" + phrase
}

func worker(id int, ch chan string) {
	message := fmt.Sprintf("Worker %d finished task", id)
	ch <- message
}

func main() {

	//using different channels
	done1 := make(chan string) // a channel is created to communicate between goroutines
	done2 := make(chan string)

	go greet("Welcome to the jungle!", done1) // go keyword is used to run the function in a separate goroutine
	go slowGreet("How ... are ... you ...?", done2)

	fmt.Println(<-done1)
	fmt.Println(<-done2)

	//Using the same channel to process multiple goroutines
	ch := make(chan string)
	numWorkers := 5

	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch)
	}

	for i := 1; i <= numWorkers; i++ {
		fmt.Println(<-ch)
	}

	close(ch)
}
