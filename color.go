package palette

import (
	"fmt"
	"strconv"
	"strings"
)

func getRGB(hex string) string {
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

func parseColor(str string) string {
	l := len(str)
	if l == 0 {
		return ""
	}

	var buf, ext = str, ""

	if idx := strings.LastIndex(str, ":"); idx != -1 {
		if idx+1 >= l {
			return ""
		}

		buf = str[idx+1:]

		if idx-2 >= 0 {
			switch str[idx-2 : idx] {
			case F8:
				ext = "38;"
			case B8:
				ext = "48;"
			}
		}

		if ext != "" {

			if _, err := strconv.Atoi(buf); err == nil && buf[0] != '#' {
				return ext + "5;" + buf
			}

			if rgb := getRGB(buf[1:]); rgb != "" {
				return ext + "2;" + rgb
			}

			return ""
		}
	}

	if l = len(buf); l == 2 && isBaseColor(buf) {
		return buf
	}

	return ""
}
