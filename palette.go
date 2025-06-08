package palette

import (
	"strings"
)

var palette *Palette

// IPalette interface
type IPalette interface {
	Use(buf string) string
	Set(k, v string) bool
	Exists(k string) bool
	Len() int
	Remove(k string)
	Reset()
}

// Palette implements IPalette
type Palette struct {
	container map[string]string
}

// NewPalette crating Palette instance
func NewPalette() *Palette {
	return &Palette{container: make(map[string]string)}
}

// NewPaletteFromConfig creating instance of Palette using IConfig
func NewPaletteFromConfig(config IConfig) *Palette {
	inst := NewPalette()

	if config != nil {
		for k, v := range config.Palette() {
			inst.Set(k, v)
		}
	}

	return inst
}

// Set key of seq
func (p *Palette) Set(k, v string) bool {
	var buf string

	if ctl := parseControl(v); ctl != "" {
		buf += sgr(ctl)
	}

	if color := parseColor(v); color != "" {
		buf += sgr(color)
	}

	if buf == "" {
		return false
	}

	p.container[k] = buf
	return true
}

// Exists check
func (p *Palette) Exists(k string) bool {
	_, ok := p.container[k]
	return ok
}

// Use is processing ur str with keys
func (p *Palette) Use(buf string) string {
	if buf != "" {
		for k, v := range p.container {
			buf = strings.ReplaceAll(buf, k, v)
		}
	}

	return buf
}

// Remove some of
func (p *Palette) Remove(k string) {
	delete(p.container, k)
}

// Reset is reset Palette obj
func (p *Palette) Reset() {
	p.container = make(map[string]string)
}

// Len ret count of keys
func (p *Palette) Len() int {
	return len(p.container)
}

// Init initialize Palette instance using Config
//
// Container of Palette initialized only once.
// U can re - palette.Init call if you want to complement  palette
//
// - cfg: ptr on Config (can be nil)
//
// noreturn
func Init(cfg IConfig) {
	palette = NewPaletteFromConfig(cfg)
}

// IsInit checks palette initialization
func IsInit() bool {
	return palette != nil
}

// Get palette instance
func Get() IPalette {
	return palette
}

// Set color by its key to the global palette
func Set(k, v string) bool {
	if IsInit() {
		return palette.Set(k, v)
	}

	return false
}

// Len of global palette
func Len() int {
	if IsInit() {
		return palette.Len()
	}

	return -1
}

// Use is using palette with replacement all known sequences
//
// return: string as its on palette.IsInit = false or empty / processed str - otherwise
func Use(buf string) string {
	if IsInit() {
		return palette.Use(buf)
	}

	return buf
}

// Reset global palette
func Reset() {
	if IsInit() {
		palette.Reset()
	}
}
