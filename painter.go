package color

// Painter is a generic interface to apply a style to a string.
type Painter interface {
	// paint applies the style to the argument and return it.
	paint(string) string
}
