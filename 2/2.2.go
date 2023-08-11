package main

import "fmt"

func main() {
	a := 0
	fmt.Printf("Hello %d", a)
	for a < 100 {
		if a == 50 {
			fmt.Println("hello")
		} else {
			fmt.Println("world")
		}
		a++
	}

}

func isTest(a int) (bool, int) {
	if a == 2 {
		return true, 1
	} else if a == 3 {
		return true, 3
	}

	return false, 4
}
