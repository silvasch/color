package main

import (
	"fmt"

	"github.com/silvasch/color"
)

func main() {
	green := color.Hex{Hex: "ffffff"}
	violet := color.New(&color.Blue{}, &color.Red{}, &green)
	fmt.Printf("Hello, %s\n", violet.Apply("World"))
}
