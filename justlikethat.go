package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("press enter to start the timer...")

	bufio.NewReader(os.Stdin).ReadBytes('\n')

	fmt.Println("timer has started!")
	timer := time.NewTimer(3 * time.Second)

	<- timer.C

	fmt.Println("timer has ended")
}