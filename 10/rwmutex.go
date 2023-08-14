package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex // Use sync.Mutex for exclusive locking
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c[key]++
}

func (c *Counter) Value(key string) int {
	c.mu.Lock() // Use Lock here to ensure consistent access
	defer c.mu.Unlock()
	return c.c[key]
}

func main() {
	key := "test"
	c := Counter{c: make(map[string]int)}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc(key)
		}()
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println(c.Value(key))
}
