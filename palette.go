package palette

import (
	"fmt"
	"strings"
)

var palette *Palette

// IPalette interface
type IPalette interface {
	Set(k, v string)
	Exists(k string) bool
	Remove(k string)
}

// Palette implements IPalette
type Palette struct {
	container map[string]string
}

// Set key of seq
func (p *Palette) Set(k, v string) {
	p.container[k] = v
}

// Exists check
func (p *Palette) Exists(k string) bool {
	_, ok := p.container[k]
	return ok
}

// Remove some of
func (p *Palette) Remove(k string) {
	delete(p.container, k)
}

// Init initialize Palette instance using Config
//
// Container of Palette initialized only once.
// U can re - palette.Init call if you want to complement  palette
//
// - cfg: ptr on Config (can be nil)
//
// noreturn
func Init(cfg *Config) {

	if palette == nil {
		palette = &Palette{make(map[string]string)}
	}

	if cfg.Palette == nil {
		return
	}

	var color, ctl string
	var buf strings.Builder
	for k, v := range cfg.Palette {
		ctl, color = "", ""
		buf.Reset()

		if ctl = parseControl(v); ctl != "" {
			if _, err := buf.WriteString(fmt.Sprintf("\x1B\x5B%s\x6d", ctl)); err != nil {
				continue
			}
		}

		if color = parseColor(v); color != "" {
			if _, err := buf.WriteString(fmt.Sprintf("\x1B\x5B%s\x6d", color)); err != nil {
				continue
			}
		}

		palette.Set(k, buf.String())
	}
}

// IsInit checks palette initialization
func IsInit() bool {
	return palette != nil
}

// Get palette instance
func Get() IPalette {
	return palette
}

// Use is using palette with replacement all known sequences
//
// return: string as its on palette.IsInit = false or empty / processed str - otherwise
func Use(buf string) string {
	if !IsInit() || buf == "" {
		return buf
	}

	for k, v := range palette.container {
		buf = strings.ReplaceAll(buf, k, v)
	}

	return buf
}
