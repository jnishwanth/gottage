package art

import (
	"github.com/lucasb-eyer/go-colorful"
)

// CharacterState defines current animation action of the 8-bit focus buddy.
type CharacterState int

const (
	StateTyping CharacterState = iota
	StateWriting
	StateSipping
)

// FocusBuddy manages 8-bit body-double companion animations.
type FocusBuddy struct {
	SkinCol    colorful.Color
	HairCol    colorful.Color
	HoodieCol  colorful.Color
	PantsCol   colorful.Color
	DeskCol    colorful.Color
}

// NewFocusBuddy creates a new 8-bit focus buddy companion.
func NewFocusBuddy() *FocusBuddy {
	return &FocusBuddy{
		SkinCol:   PicoPeach,
		HairCol:   PicoBrown,
		HoodieCol: PicoDarkPurple,
		PantsCol:  PicoNavy,
		DeskCol:   PicoBrown,
	}
}

// Draw renders the animated 8-bit focus companion at (cx, cy) on the canvas.
func (fb *FocusBuddy) Draw(c *Canvas, cx, cy int, frameCount uint64) {
	// Determine animation state based on frame cycle (sip tea every ~15 seconds / 250 frames)
	cycle := frameCount % 240
	var state CharacterState

	if cycle >= 180 && cycle < 210 {
		state = StateSipping
	} else if (cycle/30)%2 == 1 {
		state = StateWriting
	} else {
		state = StateTyping
	}

	// Gentle Breathing offset (shifts 1px up/down every 16 frames)
	breathY := 0
	if (frameCount/16)%2 == 1 {
		breathY = -1
	}

	// 1. Chair / Stool
	c.FillRect(cx+3, cy+10, 8, 8, PicoBlack)
	c.FillRect(cx+2, cy+14, 2, 8, PicoDarkGrey)
	c.FillRect(cx+10, cy+14, 2, 8, PicoDarkGrey)

	// 2. Body / Hoodie (Sitting at desk)
	c.FillRect(cx+2, cy+4+breathY, 10, 10, fb.HoodieCol)
	// Hoodie outline
	c.DrawRectOutline(cx+2, cy+4+breathY, 10, 10, PicoBlack)
	// Hood detail
	c.FillRect(cx+4, cy+4+breathY, 6, 2, PicoNavy)

	// 3. Head & Hair
	c.FillRect(cx+4, cy-2+breathY, 7, 6, fb.SkinCol)     // Head
	c.FillRect(cx+3, cy-4+breathY, 9, 3, fb.HairCol)     // Hair top
	c.FillRect(cx+3, cy-2+breathY, 3, 4, fb.HairCol)     // Hair side
	c.SetPixel(cx+8, cy+breathY, PicoBlack)              // Eye (focused down)

	// 4. State-dependent Arms & Actions
	switch state {
	case StateTyping:
		// Arms extended towards glowing keyboard
		armTick := (frameCount / 4) % 2
		if armTick == 0 {
			c.FillRect(cx+8, cy+7+breathY, 5, 2, fb.HoodieCol)
			c.FillRect(cx+12, cy+8+breathY, 3, 2, fb.SkinCol) // Hands on keyboard
		} else {
			c.FillRect(cx+8, cy+6+breathY, 5, 2, fb.HoodieCol)
			c.FillRect(cx+12, cy+7+breathY, 3, 2, fb.SkinCol)
		}

	case StateWriting:
		// Arm holding pen writing in notebook
		writeTick := (frameCount / 6) % 2
		c.FillRect(cx+8, cy+7+breathY, 4, 2, fb.HoodieCol)
		c.FillRect(cx+11+int(writeTick), cy+8+breathY, 3, 2, fb.SkinCol)
		c.SetPixel(cx+13+int(writeTick), cy+8+breathY, PicoYellow) // Pen

	case StateSipping:
		// Arm lifting mug to mouth
		c.FillRect(cx+7, cy+2+breathY, 3, 4, fb.HoodieCol)
		c.FillRect(cx+8, cy+1+breathY, 2, 2, fb.SkinCol)
		c.FillRect(cx+9, cy+breathY, 3, 3, PicoRed) // Mug raised to lips
		c.SetPixel(cx+9, cy+breathY, PicoWhite)
	}
}
