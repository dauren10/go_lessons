package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	done := make(chan bool)

	for i := 0; i < 100; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("Onle one")
			})
			fmt.Println("Multi one")
			done <- true
		}()
	}

	for i := 0; i < 100; i++ {
		<-done
	}
}
