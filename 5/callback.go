package main

import "fmt"

func doSomething(callback func(int, int) int, s string) int {
	result := callback(5, 1)
	fmt.Println(s)
	return result
}

func main() {
	someCallback := func(n1, n2 int) int {
		return n1 + n2
	}

	result := doSomething(someCallback, "hey")
	fmt.Println(result)

	multipleCallback := func(n1, n2 int) int {
		return n1 + n2
	}

	result = doSomething(multipleCallback, "mu;ti")
}
