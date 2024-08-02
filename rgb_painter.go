package color

import "fmt"

type RGB struct {
	R uint8
	G uint8
	B uint8
}

func (rgb *RGB) paint(value string) string {
	rgb_string := fmt.Sprintf("%d, %d, %d", rgb.R, rgb.G, rgb.B)
	return fmt.Sprintf("[%s] %s [/%s]", rgb_string, value, rgb_string)
}
