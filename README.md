[![Test & Build](https://github.com/rejchev/palette-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/rejchev/palette-go/actions/workflows/ci.yml)

# Palette
A simple palette based on SGR:
- [ECMA-35: Character Code Structure and Extension Techniques (eq. ISO 2022)](https://putty.org.ru/specs/ecma-035.pdf)

The Palette allows you to replace keywords in plain text with SGR before output

## Configure
First of all u must configure palette using file like `palette.json`

### Palette
`palette.json` contains one key - `palette` and this is map, where:
- `key` is user based uniq sequence
- `value` is a template sequence (ordered) that can contain:
  1. `<CTL>` аs control (predefined in `controls.go`) sequence
  2. `<F8/B8>` if u mean using 8-bit codes
  3. `<COLOR>` color vals: 
    - `30-37` (Basic Foreground) 
    - `40-47` (Basic Background) 
    - `0-255` (8-bit codes table) 
    - `#FFF000` hexed rgb (without alpha, 8-bit mode)

All of `value` parts must be delimited by `:` \
Examples: 
- `I:F8:#FFF000` indicates that i want to use italic (`I`) hexed foreground color `F8:#FFF000`
- `R` indicates that i want to reset
- `I:32` indicates that i want to use italic (`I`) base foreground color - `Green` 
- `B8:#FFFFFF` indicates that i want to use hexed background color `B8:#FFFFFF`
- `U:F8:230` indicates that i want to use underlined (`U`) [tabled](https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit) foreground color `F8:230`

note: any of controls seq is independent; they formed into a separate sgr

### Configuration definition ways

#### Directly definition
Directly definition like:
```go
package palette

import "github.com/rejchev/palette"

func getPaletteConfig() *palette.Config {
  return &palette.Config{Palette: map[string]string{
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
```
#### FS using definition
FS using definition like:
```go
package palette

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"github.com/rejchev/palette"
)

const BaseConfigPath = "/configs/palette/palette.json"

func tryLoadConfig(confpath string) (*palette.Config, error) {
  var err error
  
  var execpath string
  if execpath, err = os.Executable(); err != nil {
    return nil, err
  }

  confpath = path.Join(path.Dir(execpath), BaseConfigPath)
  if _, err = os.Stat(confpath); os.IsNotExist(err) {
    return nil, fmt.Errorf("palette file %s must be present", confpath)
  }

  var file *os.File
  if file, err = os.Open(confpath); err != nil {
    return nil, err
  }

  defer func(file *os.File) {
    _ = file.Close()
  }(file)

  var cfg palette.Config
  err = json.NewDecoder(file).Decode(&cfg)

  return &cfg, err
}
```

note: by default `config.go` provide function `TryLoadConfig`

## Usage
### Initializing
First of all u must initialize palette once using [config](#configure) (on `main` func for example):
```go
package main

import (
	"fmt"
	"github.com/rejchev/palette"
)

func main() {
  palette.Init(getPaletteConfig())
}

func getPaletteConfig() *palette.Config {
	return &palette.Config { map[string]string { 
	    // ....
	}}
}
```

### Usaging
After successfully `Init()`, you can `Use()` it:
```go
package main

import (
	"fmt"
	"github.com/rejchev/palette"
)

func main() {
  palette.Init(getPaletteConfig())
  
  fmt.Printf(palette.Use("{IF8200}[palette]: {DEF}%s\n"), "Hello, World!")
}

func getPaletteConfig() *palette.Config {
	return &palette.Config { map[string]string { 
	    // using defs from palette.json
	}}
}
```
![output](./.github/images/palette-go.png)