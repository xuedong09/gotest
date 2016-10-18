package main

import (
	. "fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i< 10; i++ {

		wg.Add(1)

		go func(id int) {

			defer wg.Done()
			time.Sleep(time.Second * 3)
			Println("goroutine", id, "done")
		}(i)
	}

	Println("main...")
	wg.Wait()
	Println("main exit.")
}
