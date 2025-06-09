package palette

import "strings"

var palette *Palette

type Entry struct {
	key, val string
}

func CreateEntry(key, val string) Entry {
	return Entry{key, val}
}

// IPalette interface
type IPalette interface {
	Use(buf string) string
	Set(k, v string)
	Exists(k string) bool
	Len() int
	Remove(k string)
	Reset()
}

// Palette implements IPalette
type Palette struct {
	container []Entry
}

// NewPalette crating Palette instance
func NewPalette() *Palette {
	return &Palette{container: make([]Entry, 0, 4)}
}

// NewPaletteFromConfig creating instance of Palette using IConfig
func NewPaletteFromConfig(config IConfig) *Palette {
	inst := NewPalette()

	if config != nil {
		for _, v := range config.Palette() {
			inst.Set(v.key, v.val)
		}
	}

	return inst
}

// Set key of seq
func (p *Palette) Set(k, v string) {
	if buf := process(v); buf != "" {
		if idx := p.Find(k); idx != -1 {
			p.container[idx].val = v
		} else {
			p.container = append(p.container, Entry{k, buf})
		}
	}
}

// Exists check
func (p *Palette) Exists(k string) bool {
	return p.Find(k) != -1
}

// Use is processing ur str with keys
func (p *Palette) Use(buf string) string {
	if buf != "" {
		for _, v := range p.container {
			buf = strings.ReplaceAll(buf, v.key, v.val)
		}
	}

	return buf
}

// Remove some of
func (p *Palette) Remove(k string) {
	if idx := p.Find(k); idx != -1 {
		p.container = append(p.container[:idx], p.container[idx+1:]...)
	}
}

// Reset is reset Palette obj
func (p *Palette) Reset() {
	p.container = make([]Entry, 0, 4)
}

// Len ret count of keys
func (p *Palette) Len() int {
	return len(p.container)
}

func (p *Palette) Find(k string) int {
	for idx, entry := range p.container {
		if entry.key == k {
			return idx
		}
	}

	return -1
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
func Set(k, v string) {
	if IsInit() {
		palette.Set(k, v)
	}
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
