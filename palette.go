package palette

import (
	"fmt"
	"strings"
)

var palette *Palette

type IPalette interface {
	Set(k, v string)
	Exists(k string) bool
	Remove(k string)
}

type Palette struct {
	container map[string]string
}

func (p *Palette) Set(k, v string) {
	p.container[k] = v
}

func (p *Palette) Exists(k string) bool {
	_, ok := p.container[k]
	return ok
}

func (p *Palette) Remove(k string) {
	delete(p.container, k)
}

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

func IsInit() bool {
	return palette != nil
}

func Get() IPalette {
	return palette
}

func Use(buf string) string {
	if !IsInit() || buf == "" {
		return buf
	}

	for k, v := range palette.container {
		buf = strings.ReplaceAll(buf, k, v)
	}

	return buf
}
