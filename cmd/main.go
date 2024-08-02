package main

import (
	"fmt"

	"github.com/silvasch/color"
)

func main() {
	black := color.New(&color.Magenta{}, &color.BgGreen{}, &color.Strikethrough{})
	fmt.Printf("%s, Hello\n", black.Apply("World"))
}
