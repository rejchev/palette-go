package palette

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// BaseConfigPath is base palette path
const BaseConfigPath = "/configs/palette/palette.json"

// Config contains map of colors
type Config struct {
	Palette map[string]string `json:"palette"`
}

// TryLoadConfig is trying load config with colors defs
//
// - confpath: str path to config (if empty - <execDir> + BaseConfigPath is used)
//
// return cfg ptr on success / err != nil on fail
func TryLoadConfig(confpath string) (*Config, error) {
	var err error
	if confpath == "" {
		var execpath string
		if execpath, err = os.Executable(); err != nil {
			return nil, err
		}

		confpath = path.Join(path.Dir(execpath), BaseConfigPath)
		if _, err = os.Stat(confpath); os.IsNotExist(err) {
			return nil, fmt.Errorf("palette file %s must be present", confpath)
		}
	}

	var file *os.File
	if file, err = os.Open(confpath); err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var config Config
	err = json.NewDecoder(file).Decode(&config)

	return &config, err
}
