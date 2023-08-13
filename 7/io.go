package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Helll owold")
	arr := make([]byte, 8)

	for {
		n, err := r.Read(arr)
		fmt.Printf("%q", arr[:n])
		fmt.Println(n, err, arr)
		fmt.Println(arr[:n])
		if err == io.EOF {
			break
		}
	}
}
