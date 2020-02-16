package main

import (
	"fmt"
)

func main() {
	N := 100

	s := fmt.Sprintf("%0128d", N)

	fmt.Println(s)
	fmt.Println(len(s))
}
