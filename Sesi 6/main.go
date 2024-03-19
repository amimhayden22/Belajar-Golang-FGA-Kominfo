package main

import (
	// "fmt"
	"sync"
	"time"
	// "runtime"
)

func main() {
	// Untuk memeriksa core di dalam laptop kita
	// fmt.Println(runtime.NumCPU())

	fruits := []string{
		"Banana",
		"Manggo",
		"Apple",
		"Watermelon",
	}

	_ = fruits

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	_, _ = wg, mutex

	var num int = 0

	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		num++
		// wg.Add(1)
		// go func() {
		// 	mutex.Lock()
		// 	num++
		// 	mutex.Unlock()
		// 	wg.Done()
		// }()
	}

	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		num--
		// wg.Add(1)
		// go func() {
		// 	mutex.Lock()
		// 	num--
		// 	mutex.Unlock()
		// 	wg.Done()
		// }()
	}

	// fmt.Printf("num => %d\n", num)

	// for index, fruit := range fruits {
	// 	i := index
	// 	f := fruit
	// 	go func() {
	// 		fmt.Printf("index: (%d), %s\n", i+1, f)
	// 		wg.Done()
	// 	}()
	// }

	wg.Wait()

	// go func ()  {
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Println("Start go routine")
	// 	wg.Done()
	// }()

	// fmt.Println("Hello World!")

	// wg.Wait()
	// time.Sleep(8 * time.Second)
}