package palette

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

var palette Palette

type IPalette interface {
	Container() map[string]string
}

type Palette struct {
	container map[string]string
}

func createPalette(cfg *Config) {
	palette = Palette{make(map[string]string)}

	var buf strings.Builder
	var color, ctl string
	// key must patterned as {<controlKey>;<F/B>;<colorName>}
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

		palette.container[k] = buf.String()
	}
}

func (p *Palette) Container() map[string]string {
	return p.container
}

func Init(cfg *Config) (err error) {
	if cfg == nil {
		var execPath string
		if execPath, err = os.Executable(); err != nil {
			return err
		}

		fpath := path.Join(path.Dir(execPath), configDir, BaseConfigFile)
		if _, err = os.Stat(fpath); os.IsNotExist(err) {
			return fmt.Errorf("palette file %s must be present", fpath)
		}

		var file *os.File
		if file, err = os.Open(fpath); err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		var config Config
		if err = json.NewDecoder(file).Decode(&config); err != nil {
			return err
		}

		cfg = &config
	}

	createPalette(cfg)

	return nil
}

func Get() IPalette {
	return &palette
}

func Use(buf string) string {
	if buf == "" {
		return buf
	}

	for k, v := range palette.Container() {
		buf = strings.ReplaceAll(buf, k, v)
	}

	return buf
}
