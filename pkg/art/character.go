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

// Draw renders the animated 16-bit focus companion at (cx, cy) on the canvas.
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
	c.FillRect(cx+2, cy+8, 8, 8, PicoBlack)
	c.FillRect(cx+1, cy+12, 2, 8, PicoDarkGrey)
	c.FillRect(cx+9, cy+12, 2, 8, PicoDarkGrey)

	// 2. Body / Cozy Hoodie (Sitting at desk)
	c.FillRect(cx+2, cy+4+breathY, 10, 10, fb.HoodieCol)
	c.DrawRectOutline(cx+2, cy+4+breathY, 10, 10, PicoBlack)
	c.FillRect(cx+4, cy+4+breathY, 6, 2, PicoNavy) // Hood collar

	// 3. Head & Hair (Pixel details)
	c.FillRect(cx+4, cy-2+breathY, 7, 6, fb.SkinCol) // Face
	c.FillRect(cx+3, cy-5+breathY, 9, 4, fb.HairCol) // Hair top
	c.FillRect(cx+3, cy-2+breathY, 2, 4, fb.HairCol) // Hair side
	c.FillRect(cx+8, cy+breathY, 1, 1, PicoBlack)    // Eye focused on work

	// 4. State-dependent Arms & Actions
	switch state {
	case StateTyping:
		armTick := (frameCount / 4) % 2
		if armTick == 0 {
			c.FillRect(cx+8, cy+7+breathY, 5, 2, fb.HoodieCol)
			c.FillRect(cx+12, cy+8+breathY, 2, 2, fb.SkinCol) // Hands on keyboard
		} else {
			c.FillRect(cx+8, cy+6+breathY, 5, 2, fb.HoodieCol)
			c.FillRect(cx+12, cy+7+breathY, 2, 2, fb.SkinCol)
		}

	case StateWriting:
		writeTick := (frameCount / 6) % 2
		c.FillRect(cx+8, cy+7+breathY, 4, 2, fb.HoodieCol)
		c.FillRect(cx+11+int(writeTick), cy+8+breathY, 2, 2, fb.SkinCol)
		c.FillRect(cx+13+int(writeTick), cy+8+breathY, 1, 1, PicoYellow) // Pen

	case StateSipping:
		c.FillRect(cx+7, cy+2+breathY, 3, 4, fb.HoodieCol)
		c.FillRect(cx+9, cy+1+breathY, 2, 2, fb.SkinCol)
		c.FillRect(cx+10, cy+breathY, 3, 3, PicoRed) // Red lofi mug to lips
		c.SetPixel(cx+10, cy+breathY, PicoWhite)
	}
}
