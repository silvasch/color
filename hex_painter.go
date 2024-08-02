package color

import (
	"log"
	"strconv"
	"strings"
)

type Hex struct {
	Hex string
}

func (hex *Hex) paint(value string) string {
	if len(hex.Hex) != 6 {
		log.Fatal("hex string has to contain exactly six characters")
	}

	chars := strings.Split(hex.Hex, "")

	r_string := chars[0] + chars[1]
	g_string := chars[2] + chars[3]
	b_string := chars[4] + chars[5]

	r, err := strconv.ParseUint(r_string, 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	g, err := strconv.ParseUint(g_string, 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	b, err := strconv.ParseUint(b_string, 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	color := RGB{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
	}

	return color.paint(value)
}
