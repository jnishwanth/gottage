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

// FocusBuddy manages 16-bit body-double companion animations.
type FocusBuddy struct {
	SkinCol   colorful.Color
	HairCol   colorful.Color
	HoodieCol colorful.Color
	PantsCol  colorful.Color
	DeskCol   colorful.Color
}

// NewFocusBuddy creates a new 16-bit focus buddy companion.
func NewFocusBuddy() *FocusBuddy {
	return &FocusBuddy{
		SkinCol:   PicoPeach,
		HairCol:   PicoBrown,
		HoodieCol: PicoDarkPurple,
		PantsCol:  PicoNavy,
		DeskCol:   PicoBrown,
	}
}

// Draw renders the animated 16-bit focus companion at (cx, cy) on the 160x90 canvas.
func (fb *FocusBuddy) Draw(c *Canvas, cx, cy int, frameCount uint64) {
	cycle := frameCount % 240
	var state CharacterState

	if cycle >= 180 && cycle < 210 {
		state = StateSipping
	} else if (cycle/30)%2 == 1 {
		state = StateWriting
	} else {
		state = StateTyping
	}

	// Gentle Breathing offset
	breathY := 0
	if (frameCount/16)%2 == 1 {
		breathY = -1
	}

	// 1. Chair & Backrest
	c.FillRect(cx+6, cy+18, 16, 18, PicoBlack)
	c.FillRect(cx+4, cy+26, 4, 18, PicoDarkGrey)
	c.FillRect(cx+20, cy+26, 4, 18, PicoDarkGrey)

	// 2. Body / Cozy Hoodie (Sitting at desk)
	c.FillRect(cx+4, cy+8+breathY, 20, 20, fb.HoodieCol)
	c.DrawRectOutline(cx+4, cy+8+breathY, 20, 20, PicoBlack)
	c.FillRect(cx+8, cy+8+breathY, 12, 4, PicoNavy) // Hood collar

	// 3. Head & Hair (HD Pixel details)
	c.FillRect(cx+8, cy-4+breathY, 14, 12, fb.SkinCol)  // Face
	c.FillRect(cx+6, cy-8+breathY, 18, 6, fb.HairCol)   // Hair top
	c.FillRect(cx+6, cy-4+breathY, 6, 8, fb.HairCol)    // Hair side
	c.FillRect(cx+16, cy+breathY, 2, 2, PicoBlack)      // Eye focused on work

	// 4. State-dependent Arms & Actions
	switch state {
	case StateTyping:
		armTick := (frameCount / 4) % 2
		if armTick == 0 {
			c.FillRect(cx+16, cy+14+breathY, 10, 4, fb.HoodieCol)
			c.FillRect(cx+24, cy+16+breathY, 6, 4, fb.SkinCol) // Hands on keyboard
		} else {
			c.FillRect(cx+16, cy+12+breathY, 10, 4, fb.HoodieCol)
			c.FillRect(cx+24, cy+14+breathY, 6, 4, fb.SkinCol)
		}

	case StateWriting:
		writeTick := (frameCount / 6) % 2
		c.FillRect(cx+16, cy+14+breathY, 8, 4, fb.HoodieCol)
		c.FillRect(cx+22+int(writeTick)*2, cy+16+breathY, 6, 4, fb.SkinCol)
		c.FillRect(cx+26+int(writeTick)*2, cy+16+breathY, 2, 2, PicoYellow) // Pen

	case StateSipping:
		c.FillRect(cx+14, cy+4+breathY, 6, 8, fb.HoodieCol)
		c.FillRect(cx+16, cy+2+breathY, 4, 4, fb.SkinCol)
		c.FillRect(cx+18, cy+breathY, 6, 6, PicoRed) // Red lofi mug to lips
		c.SetPixel(cx+18, cy+breathY, PicoWhite)
	}
}
