package core

import (
	"image/color"
)

type Color32 struct {
	color.RGBA `xml:"color,attr,omitempty"`
}

func (c Color32) Colored32() (R, G, B, A float32) {
	return float32(c.R), float32(c.G), float32(c.B), float32(c.A)
}

func (c Color32) Colored64() (R, G, B, A float64) {
	return float64(c.R), float64(c.G), float64(c.B), float64(c.A)
}

type Color64 struct {
	color.RGBA64 `xml:"color,attr,omitempty"`
}

func (c Color64) Colored32() (R, G, B, A float32) {
	return float32(c.R), float32(c.G), float32(c.B), float32(c.A)
}

func (c Color64) Colored64() (R, G, B, A float64) {
	return float64(c.R), float64(c.G), float64(c.B), float64(c.A)
}

type Colored interface {
	Colored32() (R, G, B, A float32)
	Colored64() (R, G, B, A float64)
}
