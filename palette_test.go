package palette

import (
	"fmt"
	"testing"
)

func loadPaletteConfig() *Config {
	if cfg, err := TryLoadConfig(""); err == nil {
		return cfg
	}

	return nil
}

func getPaletteConfig() *Config {
	return &Config{Palette: map[string]string{
		"{DEF}":      "R",
		"{FBA}":      "30",
		"{BBU}":      "44",
		"{BMA}":      "45",
		"{BCY}":      "46",
		"{BWH}":      "47",
		"{F8BA}":     "F8:0",
		"{F8RD}":     "F8:1",
		"{F8GR}":     "F8:2",
		"{F8YL}":     "F8:3",
		"{F8BU}":     "F8:4",
		"{F8MA}":     "F8:5",
		"{F8CY}":     "F8:6",
		"{F8WH}":     "F8:7",
		"{IF8200}":   "I:F8:200",
		"{IF8APPLE}": "I:F8:#FFF000",
	}}
}

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

func TestInit(t *testing.T) {
	Init(getPaletteConfig())

	if !IsInit() || !Get().Exists("{IF8APPLE}") {
		t.Error("Expected exists = {IF8APPLE}")
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

	Init(getPaletteConfig())

	src := "{FBA}[palette]: {DEF}Hello, world!"
	expected := fmt.Sprintf("\u001B[30m[palette]: \u001B[0mHello, world!")

	if val := Use(src); val != expected {
		t.Errorf("Use() = %s, want %s", val, expected)
	}
}
