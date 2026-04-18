package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	inputChannel := make(chan string)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func()  {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			inputChannel <- text
		}
	}()

	fmt.Println("waiting for new input every 10 seconds....")

	for {
		select {
		case <- ticker.C :
			fmt.Println("Tick! input check")
			select {
			case val := <- inputChannel :
				// input := strings.TrimSuffix(val, "\n")
				fmt.Printf("received %s", val)
			default:
				fmt.Println("no new input found.")
			}
		}
	}
}