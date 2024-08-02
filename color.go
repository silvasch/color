package color

// Color is used to apply single or combinded painters to strings.
type Color struct {
	painters []Painter
}

// New creates a new Color with the given Painters.
func New(painters ...Painter) Color {
	color := Color{}

	for _, colorer := range painters {
		color.painters = append(color.painters, colorer)
	}

	return color
}

// Apply applies the style of the Color to the given string and returns a styled version.
func (c Color) Apply(value string) string {
	for _, painter := range c.painters {
		value = painter.paint(value)
	}
	return value
}
