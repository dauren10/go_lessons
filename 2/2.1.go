package main

import "fmt"

func main() {
	//sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	sum := 0
	for sum < 1000 {
		sum += 100
	}
	fmt.Println(sum)
	k := 0
	for k < 100 {
		k += 10
		fmt.Println(k)
	}

	for true {
		fmt.Println(11)
	}

	for {
		fmt.Println(11)
	}
}
