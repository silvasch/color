package color

import "fmt"

type Red struct{}

func (r *Red) paint(value string) string {
	return fmt.Sprintf("[red] %v [/red]", value)
}

type Green struct{}

func (g *Green) paint(value string) string {
	return fmt.Sprintf("[green] %v [/green]", value)
}

type Blue struct{}

func (g *Blue) paint(value string) string {
	return fmt.Sprintf("[blue] %v [/blue]", value)
}
