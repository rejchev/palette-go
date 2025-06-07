# Palette
A simple palette based on SGR:
- [ECMA-35: Character Code Structure and Extension Techniques (eq. ISO 2022)](https://putty.org.ru/specs/ecma-035.pdf)

The Palette allows you to replace keywords in plain text with SGR before output

## Usage
### Initializing
First of all u must initialize palette once (on `main` func for example):

```go

import "github.com/rejchev/palette"

func main() {
	if err := palette.Init(nil); err != nil {
		fmt.Printf("Init() error = %v\n", err)
	}
}
```
Function `Init(.)` takes one parameter of `*palette.Config` type with palette.

### Configuring
By default config must be existing on `<executerDir>/configs/palette/palette.json` path \
You can override it before `Init()`:
```go
import "github.com/rejchev/palette"

func main() {
  // <executerDir>/../configs/palette
  palette.SetConfigDirectory("../configs/palette")
  if err := palette.Init(nil); err != nil {
		fmt.Printf("Init() error = %v\n", err)
  }
}
```
By the way you can provide config directly:
```go
import "github.com/rejchev/palette"

func main() {
	if err := palette.Init(getPaletteConfig()); err != nil {
		fmt.Printf("Init() error = %v\n", err)
	}
}

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
### Usaging
After successfully `Init()`, you can `Use()` it:
```go
import (
	"fmt"
	"github.com/rejchev/palette"
)

func main() {
	if err := palette.Init(getPaletteConfig()); err != nil {
		fmt.Printf("Init() error = %v", err)
	}

	fmt.Printf(palette.Use("{IF8200}[palette]: {DEF}%s\n"), "Hello, World")
}

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
![output](./.github/images/palette-go.png)
