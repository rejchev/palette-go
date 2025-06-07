package palette

const BaseConfigFile = "palette.json"

var configDir = "configs/palette"

type Config struct {
	Palette map[string]string `json:"palette"`
}

func SetConfigDirectory(dir string) {
	configDir = dir
}

func GetConfigDirectory() string {
	return configDir
}
