package color

import "fmt"

type RGB struct {
	R uint8
	G uint8
	B uint8
}

func (rgb *RGB) paint(value string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", rgb.R, rgb.G, rgb.B, value)
}
