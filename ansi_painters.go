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

type BgBlack struct{}

func (c *BgBlack) paint(value string) string {
	return fmt.Sprintf("\033[40m%s\033[49m", value)
}

type BgRed struct{}

func (c *BgRed) paint(value string) string {
	return fmt.Sprintf("\033[41m%s\033[49m", value)
}

type BgGreen struct{}

func (c *BgGreen) paint(value string) string {
	return fmt.Sprintf("\033[42m%s\033[49m", value)
}

type BgYellow struct{}

func (c *BgYellow) paint(value string) string {
	return fmt.Sprintf("\033[43m%s\033[49m", value)
}

type BgBlue struct{}

func (c *BgBlue) paint(value string) string {
	return fmt.Sprintf("\033[44m%s\033[49m", value)
}

type BgMagenta struct{}

func (c *BgMagenta) paint(value string) string {
	return fmt.Sprintf("\033[45m%s\033[49m", value)
}

type BgCyan struct{}

func (c *BgCyan) paint(value string) string {
	return fmt.Sprintf("\033[46m%s\033[49m", value)
}

type BgWhite struct{}

func (c *BgWhite) paint(value string) string {
	return fmt.Sprintf("\033[47m%s\033[49m", value)
}
