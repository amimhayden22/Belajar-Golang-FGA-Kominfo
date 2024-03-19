package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {

		time.Sleep(3 * time.Second)

		result := <-ch
		fmt.Println("result => ", result)
	}()

	ch <- "First Data"

	fmt.Println("Done!");
}