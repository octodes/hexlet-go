package main

import (
	"fmt"

	"github.com/octodes/hexlet-go/greeting"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(greeting.Hello())

	color.Cyan(greeting.Hello())
}
