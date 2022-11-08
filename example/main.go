package main

import (
	"fmt"
	termcolor "go-term-colorize"
)

func main() {
	str := "hello world"
	str = termcolor.BrightWhite.Foreground(str)
	fmt.Println(str + termcolor.Reset)

	str = "world hello"
	str = termcolor.MakeColorCode(255, 255, 255).Foreground(str)
	fmt.Println(str + termcolor.Reset)
}
