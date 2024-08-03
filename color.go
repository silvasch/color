package color

import (
	"fmt"
	"io"
)

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

// Fprint is the same as fmt.Fprint, but it also applies the color.
func (c Color) Fprint(w io.Writer, a ...any) (int, error) {
	return fmt.Fprint(w, c.Sprint(a...))
}

// Fprintf is the same as fmt.Fprintf, but it also applies the color.
func (c Color) Fprintf(w io.Writer, format string, a ...any) (int, error) {
	return fmt.Fprint(w, c.Sprintf(format, a...))
}

// Fprintln is the same as fmt.Fprintln, but it also applies the color.
func (c Color) Fprintln(w io.Writer, a ...any) (int, error) {
	return fmt.Fprintln(w, c.Sprint(a...))
}

// Print is the same as fmt.Print, but it also applies the color.
func (c Color) Print(a ...any) (int, error) {
	return fmt.Print(c.Sprint(a...))
}

// Printf is the same as fmt.Printf, but it also applies the color.
func (c Color) Printf(format string, a ...any) (int, error) {
	return fmt.Print(c.Sprintf(format, a...))
}

// Println is the same as fmt.Println, but it also applies the color.
func (c Color) Println(a ...any) (int, error) {
	return fmt.Println(c.Sprint(a...))
}

// Sprint is the same as fmt.Sprint, but it also applies the color.
func (c Color) Sprint(a ...any) string {
	return c.apply(fmt.Sprint(a...))
}

// Sprintf is the same as fmt.Sprintf, but it also applies the color.
func (c Color) Sprintf(format string, a ...any) string {
	return c.apply(fmt.Sprintf(format, a...))
}

// Sprintln is the same as fmt.Sprintln, but it also applies the color.
func (c Color) Sprintln(a ...any) string {
	return c.apply(fmt.Sprintln(a...))
}

func (c Color) apply(value string) string {
	for _, painter := range c.painters {
		value = painter.paint(value)
	}
	return value
}
