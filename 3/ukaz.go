package main

import "fmt"

type Point struct {
	X int //с большой отвечает за область видимости переменных
	Y int
}

func main() {
	pointers()
}

// в golang нет классов
func pointers() {
	a := "hello world"

	fmt.Println(a)
	p := &a
	fmt.Println(p)  //ssylka byte
	fmt.Println(*p) //znachenie po ssylke
	*p = "dauren"   //изменили значение а через ссылку на переменную а
	fmt.Println(a)

	b := 42
	g := &b
	*g = *g / 2
	fmt.Println(b)
}
