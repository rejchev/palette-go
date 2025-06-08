package palette

type IConfig interface {
	Palette() map[string]string
}

// Config contains map of colors
type Config struct {
	palette map[string]string
}

func NewConfig(m map[string]string) *Config {
	return &Config{m}
}

func (c *Config) Palette() map[string]string {
	return c.palette
}

func GetBasePaletteConfig() IConfig {
	return NewConfig(map[string]string{

		// Base controls
		"{B}":   "B",   // Bold
		"{L}":   "L",   // Light
		"{I}":   "I",   // Italic
		"{R}":   "R",   // Reset
		"{U}":   "U",   // Underlined
		"{SB}":  "SB",  // Slow blink
		"{REV}": "REV", // Swap foreground and background colors; inconsistent emulation

		// Base Foreground colors
		"{FBA}": "30", // Black
		"{FRD}": "31", // Red
		"{FGR}": "32", // Green
		"{FYL}": "33", // Yellow
		"{FBU}": "34", // Blue
		"{FMA}": "35", // Magenta
		"{FCY}": "36", // Cyan
		"{FWH}": "37", // White

		// Base Background colors
		"{BBA}": "40", // Black
		"{BRD}": "41", // Red
		"{BGR}": "42", // Green
		"{BYL}": "43", // Yellow
		"{BBU}": "44", // Blue
		"{BMA}": "45", // Magenta
		"{BCY}": "46", // Cyan
		"{BWH}": "47", // White
	})
}
