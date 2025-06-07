# Palette
A simple palette based on SGR:
- [ECMA-35: Character Code Structure and Extension Techniques (eq. ISO 2022)](https://putty.org.ru/specs/ecma-035.pdf)

The Palette allows you to replace keywords in plain text with SGR before output

## Usage
### Configuring
First of all u must configure palette using file like `palette.json` \

#### Palette
`palette.json` contains one key - `palette` and this is map, there:
- `key` is user based uniq sequence
- `value` is patterned sequence that might be contains (ordered):
  1. `<CTL>` аs control (predefined in `controls.go`) sequence
  2. `<F8/B8>` if u mean using 8-bit colors
  3. `<COLOR>` color val: `30-37 (Basic Foreground)`, `40-47 (Basic Background)`, `0-255 (8-bit color table)`, `#FFFFFF 8-bit colors`

All of `value` parts must be delimited by `:` \
For example: 
- `I:F8:#FFF000` indicates that i want to use italic (`I`) hexed color `F8:#FFF000`
- `R` indicates that i want to reset
- `I:32` indicates that i want to use italic (`I`) base color - `Green` 
