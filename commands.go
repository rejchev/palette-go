package palette

import (
	"fmt"
	"strconv"
	"strings"
)

const delimiter = ":"

type commandFn = func(string, int) string
type commands = map[string]commandFn

var cmds = commands{
	// Reset or Normal
	"R": processPrimitive,
	// Bold
	"B": processPrimitive,
	// May be implemented as a light font weight like bold
	"L": processPrimitive,
	// Italic
	"I": processPrimitive,
	// Underline
	"U": processPrimitive,
	// Slow blink
	"SB": processPrimitive,
	// Rapid blink
	"RB": processPrimitive,
	// Swap foreground and background colors; inconsistent emulation
	"REV": processPrimitive,
	// Conceal or hide (Not widely supported)
	"HIDE": processPrimitive,
	// Crossed-out, or strike
	"CO": processPrimitive,
	// Any of base colors: 30-37 or 40-47 or 90-97 or 100-107
	"C": processColor,
	// 8bit foreground tabled colors
	"FTC": process8BitColor,
	// 8bit background tabled colors
	"BTC": process8BitColor,
	// 8bit underline tabled colors (not in standard; implemented in Kitty, VTE, mintty, and iTerm2)
	"UTC": process8BitColor,
	// 24bit foreground hex colors
	"FHC": process24BitColor,
	// 24bit background hex colors
	"BHC": process24BitColor,
	// 24bit underline hex colors
	"ETH": process24BitColor,
}

func processPrimitive(s string, idx int) string {
	switch s {
	case "R":
		return "0"
	case "B":
		return "1"
	case "L":
		return "2"
	case "I":
		return "3"
	case "U":
		return "4"
	case "SB":
		return "5"
	case "RB":
		return "6"
	case "REV":
		return "7"
	case "HIDE":
		return "8"
	case "CO":
		return "9"
	}
	return ""
}

func processColor(s string, idx int) string {
	if idx == -1 {
		return ""
	}

	if _, err := strconv.Atoi(s[idx+1:]); err != nil {
		return ""
	}

	val := s[idx+1:]
	switch {
	case isFBC(val) || isBBC(val) || isFBBC(val) || isBBBC(val):
		return val
	}

	return ""
}

func process8BitColor(s string, idx int) string {
	if idx == -1 {
		return ""
	}

	val := s[idx+1:]
	if _, err := strconv.Atoi(val); err != nil {
		return ""
	}

	switch s[:idx] {
	case "FTC":
		return "38;5;" + val
	case "BTC":
		return "48;5;" + val
	case "UTC":
		return "58;5;" + val
	}

	return ""
}

func process24BitColor(s string, idx int) string {
	if idx == -1 {
		return ""
	}

	val := getRgb(s[idx+2:])
	if val == "" {
		return ""
	}

	switch s[:idx] {
	case "FHC":
		return "38;2;" + val
	case "BHC":
		return "48;2;" + val
	case "UHC":
		return "58;2;" + val
	}
	return ""
}

func sgr(val string) string {

	if val == "" {
		return val
	}

	command := val
	idx := strings.Index(val, delimiter)
	if idx != -1 {
		command = val[:idx]
	}

	if fn, ok := cmds[command]; ok {
		if result := fn(val, idx); result != "" {
			return fmt.Sprintf("\x1B\x5B%s\x6D", result)
		}
	}

	return val
}

func isFBC(buf string) bool {
	return len(buf) == 2 && buf[0] == '3' && buf[1] >= '0' && buf[1] <= '7'
}

func isBBC(buf string) bool {
	return len(buf) == 2 && buf[0] == '4' && buf[1] >= '0' && buf[1] <= '7'
}

func isFBBC(buf string) bool {
	return len(buf) == 2 && buf[0] == '9' && buf[1] >= '0' && buf[1] <= '7'
}

func isBBBC(buf string) bool {
	return len(buf) == 3 && buf[0] == '1' && buf[1] == '0' && buf[2] >= '0' && buf[2] <= '7'
}

func getRgb(hex string) string {
	if len(hex) != 6 {
		return ""
	}

	if val, err := strconv.ParseUint(hex, 16, 32); err == nil {
		return fmt.Sprintf("%d;%d;%d",
			uint8(val>>16),
			uint8((val>>8)&0xFF),
			uint8(val&0xFF))
	}

	return ""
}
