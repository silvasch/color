package color

import "fmt"

type Black struct{}

func (c *Black) paint(value string) string {
	return fmt.Sprintf("\033[30m%s\033[39m", value)
}

type Red struct{}

func (c *Red) paint(value string) string {
	return fmt.Sprintf("\033[31m%s\033[39m", value)
}

type Green struct{}

func (c *Green) paint(value string) string {
	return fmt.Sprintf("\033[32m%s\033[39m", value)
}

type Yellow struct{}

func (c *Yellow) paint(value string) string {
	return fmt.Sprintf("\033[33m%s\033[39m", value)
}

type Blue struct{}

func (c *Blue) paint(value string) string {
	return fmt.Sprintf("\033[34m%s\033[39m", value)
}

type Magenta struct{}

func (c *Magenta) paint(value string) string {
	return fmt.Sprintf("\033[35m%s\033[39m", value)
}

type Cyan struct{}

func (c *Cyan) paint(value string) string {
	return fmt.Sprintf("\033[36m%s\033[39m", value)
}

type White struct{}

func (c *White) paint(value string) string {
	return fmt.Sprintf("\033[37m%s\033[39m", value)
}
