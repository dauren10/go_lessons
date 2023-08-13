package main

import "fmt"

func main() {
	var json map[string]interface{}
	var b interface{} = "3.14"
	var a interface{} = true

	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("str")
	default:
		fmt.Println("def")
	}

}
