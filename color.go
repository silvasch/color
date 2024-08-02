package color

type Color struct {
	painters []Painter
}

func New(painters ...Painter) Color {
	color := Color{}

	for _, colorer := range painters {
		color.painters = append(color.painters, colorer)
	}

	return color
}

func (c Color) Apply(value string) string {
	for _, painter := range c.painters {
		value = painter.paint(value)
	}
	return value
}
