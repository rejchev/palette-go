package palette

import (
	"fmt"
	"testing"
)

func TestParseHex(t *testing.T) {
	hex := "FFFFFF"
	expected := "255;255;255"
	if val := getRgb(hex); val != expected {
		t.Errorf("Expected %s, got %s", expected, val)
	}
}

func TestParseHexInvalid(t *testing.T) {
	hex := "#ce2626"
	expected := "206;38;38"
	if val := getRgb(hex[1:]); val != expected {
		t.Errorf("Expected %s, got %s", expected, val)
	}
}

func TestParseHex2(t *testing.T) {
	hex := "FFF"
	expected := ""
	if val := getRgb(hex); val != expected {
		t.Errorf("Expected %s, got %s", expected, val)
	}
}

func TestParseControl(t *testing.T) {
	command := "R"
	expected := processPrimitive(command, -1)

	if result := cmds[command](command, -1); result != expected {
		t.Errorf("parseControl() = %s, want %s", result, expected)
	}
}

func TestParseColor(t *testing.T) {
	color := "C:30"
	expected := "30"

	if result := processColor(color, 1); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParseColorInvalid(t *testing.T) {
	color := "C:29"
	expected := ""
	if result := processColor(color, 1); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParseColorInvalid2(t *testing.T) {
	color := "C:48"
	expected := ""
	if result := processColor(color, 1); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParseColor2(t *testing.T) {
	color := "C:47"
	expected := "47"
	if result := processColor(color, 1); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParse8Color(t *testing.T) {
	color := "FTC:1"
	expected := "38;5;1"
	if result := process8BitColor(color, 3); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParse8Color2(t *testing.T) {
	color := "BTC:200"
	expected := "48;5;200"
	if result := process8BitColor(color, 3); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParse8Color3(t *testing.T) {
	color := "FTC:200"
	expected := "38;5;200"
	if result := process8BitColor(color, 3); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParse24Color(t *testing.T) {
	color := "FHC:#FFF000"
	expected := "38;2;255;240;0"
	if result := process24BitColor(color, 3); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParse24ColorInvalid(t *testing.T) {
	//
	color := "FHC:FFF000"
	expected := ""
	if result := process24BitColor(color, 3); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestUseNil(t *testing.T) {

	Init(GetBasePaletteConfig())

	src := "{FBA}[palette]: {R}Hello, world!"
	expected := fmt.Sprintf("\u001B[30m[palette]: \u001B[0mHello, world!")

	if val := Use(src); val != expected {
		t.Errorf("Use() = %s, want %s", val, expected)
	}
}
