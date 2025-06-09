package palette

type IConfig interface {
	Palette() []Entry
}

// Config contains map of colors
type Config struct {
	palette []Entry
}

func NewConfig(m []Entry) *Config {
	return &Config{m}
}

func (c *Config) Palette() []Entry {
	return c.palette
}

func GetBasePaletteConfig() IConfig {
	return NewConfig([]Entry{
		// Base controls
		{"{R}", "R"},       // Reset
		{"{B}", "B"},       // Bold
		{"{L}", "L"},       // Light
		{"{I}", "I"},       // Italic
		{"{U}", "U"},       // Underlined
		{"{SB}", "SB"},     // Slow blink
		{"{RB}", "RB"},     // Rapid blink
		{"{REV}", "REV"},   // Swap foreground and background colors; inconsistent emulation
		{"{HIDE}", "HIDE"}, // Conceal or hide (Not widely supported)
		{"CO", "CO"},       // Crossed-out, or strike

		// Base Foreground colors
		{"{FBA}", "C:30"}, // Black
		{"{FRD}", "C:31"}, // Red
		{"{FGR}", "C:32"}, // Green
		{"{FYL}", "C:33"}, // Yellow
		{"{FBU}", "C:34"}, // Blue
		{"{FMA}", "C:35"}, // Magenta
		{"{FCY}", "C:36"}, // Cyan
		{"{FWH}", "C:37"}, // White

		// Base background colors
		{"{BBA}", "C:40"}, // Black
		{"{BRD}", "C:41"}, // Red
		{"{BGR}", "C:42"}, // Green
		{"{BYL}", "C:43"}, // Yellow
		{"{BBU}", "C:44"}, // Blue
		{"{BMA}", "C:45"}, // Magenta
		{"{BCY}", "C:46"}, // Cyan
		{"{BWH}", "C:47"}, // White
	})
}
