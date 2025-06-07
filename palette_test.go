package palette

import (
	"fmt"
	"testing"
)

func TestParseHex(t *testing.T) {
	hex := "FFFFFF"
	expected := "255;255;255"
	if val := getRGB(hex); val != expected {
		t.Errorf("Expected %s, got %s", expected, val)
	}
}

func TestParseHexInvalid(t *testing.T) {
	hex := "#ce2626"
	expected := "206;38;38"
	if val := getRGB(hex[1:]); val != expected {
		t.Errorf("Expected %s, got %s", expected, val)
	}
}

func TestParseHex2(t *testing.T) {
	hex := "FFF"
	expected := ""
	if val := getRGB(hex); val != expected {
		t.Errorf("Expected %s, got %s", expected, val)
	}
}

func TestNilInit(t *testing.T) {
	if err := Init(nil); err != nil {
		t.Errorf("Init() error = %v, want nil", err)
	}
}

func TestInit(t *testing.T) {

	cfg := Config{map[string]string{
		"{FBL}": "30",
		"{FRD}": "31",
		"{FR}":  "32",
		"{BBL}": "40",
		"{R}":   "R",
	}}

	if err := Init(&cfg); err != nil {
		t.Errorf("Init() error = %v, want nil", err)
	}
}

func TestSetConfigDirectory(t *testing.T) {
	SetConfigDirectory("test")

	if err := Init(nil); err == nil {
		t.Errorf("Init() error = nil, want not nil")
	}
}

func TestParseControl(t *testing.T) {
	ctl := "R"
	expected := controls[ctl]

	if result := parseControl(ctl); result != expected {
		t.Errorf("parseControl() = %s, want %s", result, expected)
	}
}

func TestParseUnknownControl(t *testing.T) {
	ctl := "RS"
	expected := ""

	if result := parseControl(ctl); result != expected {
		t.Errorf("parseControl() = %s, want %s", result, expected)
	}
}

func TestParseMediumControl(t *testing.T) {
	ctl := "R"
	expected := controls[ctl]

	if result := parseControl(ctl + ":31"); result != expected {
		t.Errorf("parseControl() = %s, want %s", result, expected)
	}
}

func TestParseColor(t *testing.T) {
	color := "30"
	expected := "30"

	if result := parseColor(color); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParseColorInvalid(t *testing.T) {
	color := "29"
	expected := ""
	if result := parseColor(color); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParseColorInvalid2(t *testing.T) {
	color := "48"
	expected := ""
	if result := parseColor(color); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParseColor2(t *testing.T) {
	color := ":47"
	expected := "47"
	if result := parseColor(color); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParse8Color(t *testing.T) {
	color := "F8:1"
	expected := "38;5;1"
	if result := parseColor(color); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParse8Color2(t *testing.T) {
	color := "B8:200"
	expected := "48;5;200"
	if result := parseColor(color); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParse8Color3(t *testing.T) {
	color := "I:F8:200"
	expected := "38;5;200"
	if result := parseColor(color); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParseHardControl(t *testing.T) {
	ctl := "I"
	expected := controls[ctl]

	if result := parseControl(ctl + ":F8:201"); result != expected {
		t.Errorf("parseControl() = %s, want %s", result, expected)
	}
}

func TestParse8Color4(t *testing.T) {
	// Italic foreground rgb(255;240;0)
	color := "I:F8:#FFF000"
	expected := "38;2;255;240;0"
	if result := parseColor(color); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParse8ColorInvalid(t *testing.T) {
	//
	color := "I:F9:#FFF000"
	expected := ""
	if result := parseColor(color); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestParse8ColorInvalid2(t *testing.T) {
	//
	color := "I:B8:FFF000"
	expected := ""
	if result := parseColor(color); result != expected {
		t.Errorf("parseColor() = %s, want %s", result, expected)
	}
}

func TestUseNil(t *testing.T) {
	SetConfigDirectory("configs/palette")
	if err := Init(nil); err != nil {
		t.Errorf("Init() error = %v, want nil", err)
	}

	src := "{FGR}[palette]: {DEF}Hello, world!"
	expected := fmt.Sprintf("\u001B[32m[palette]: \u001B[0mHello, world!")

	if val := Use(src); val != expected {
		t.Errorf("Use() = %s, want %s", val, expected)
	}
}
