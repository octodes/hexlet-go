package main

import (
	"fmt"

	"./greeting" // Относительный импорт
)

func main() {
	fmt.Println(greeting.Get())
}
