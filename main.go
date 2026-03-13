package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	screen := Add(2, 2)
	fmt.Println("Я домашка")
	fmt.Printf("%d", screen)
}
