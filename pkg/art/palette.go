package art

import (
	"github.com/lucasb-eyer/go-colorful"
)

// PICO-8 16-Color Retro Palette Definitions
var (
	PicoBlack      = hex("#000000") // 0: Black
	PicoNavy       = hex("#1D2B53") // 1: Dark Blue
	PicoDarkPurple = hex("#7E2553") // 2: Dark Purple
	PicoGreen      = hex("#008751") // 3: Dark Green
	PicoBrown      = hex("#AB5236") // 4: Brown
	PicoDarkGrey   = hex("#5F574F") // 5: Dark Grey
	PicoLightGrey  = hex("#C2C3C7") // 6: Light Grey
	PicoWhite      = hex("#FFF1E8") // 7: White
	PicoRed        = hex("#FF004D") // 8: Red
	PicoOrange     = hex("#FFA300") // 9: Orange
	PicoYellow     = hex("#FFEC27") // 10: Yellow
	PicoLime       = hex("#00E436") // 11: Green
	PicoCyan       = hex("#29ADFF") // 12: Blue
	PicoLavender   = hex("#83769C") // 13: Indigo / Lavender
	PicoPink       = hex("#FF77A8") // 14: Pink
	PicoPeach      = hex("#FFCCAA") // 15: Peach
)

func hex(s string) colorful.Color {
	c, _ := colorful.Hex(s)
	return c
}
