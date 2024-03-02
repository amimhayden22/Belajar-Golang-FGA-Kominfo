package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{10, 20, 30, 40, 50, 60}

	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	for _, v := range nums {
		go func(num int){
			fmt.Println(num)
			wg.Done()
		}(v)
	}

	wg.Wait()
}