package palette

import "strings"

const (
	F8 = "F8"
	B8 = "B8"
)

type Controls = map[string]string

var controls = Controls{

	// Reset or Normal
	"R": "0",

	// Bold
	"B": "1",

	// May be implemented as a light font weight like bold
	"L": "2",

	// Italic
	"I": "3",

	// Underline
	"U": "4",

	// Slow blink
	"SB": "5",

	// Swap foreground and background colors; inconsistent emulation
	"REV": "7",
}

func parseControl(s string) string {
	var idx int
	if idx = strings.Index(s, ":"); idx == -1 {
		idx = len(s)
	}

	if val, ok := controls[s[:idx]]; ok {
		return val
	}

	return ""
}

func isBaseColor(buf string) bool {
	return (buf[0] == '3' || buf[0] == '4') && (buf[1] >= '0' && buf[1] <= '7')
}
