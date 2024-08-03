package main

import (
	"github.com/silvasch/color"
)

func main() {
	color := color.New(&color.Blue{}, &color.Underline{})
	color.Printf("Hello, %s!\n", "World")
}
