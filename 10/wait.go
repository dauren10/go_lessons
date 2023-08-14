package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // n gorutin
		k := i
		go func() {
			defer wg.Done()
			fmt.Println("gorotine sleeping 300ms", k)
			time.Sleep(300 * time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Print("done...")

}
