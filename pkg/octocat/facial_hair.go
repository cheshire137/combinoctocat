package models

type FacialHair struct {
	Color *Color
	Style string
}

func NewFacialHair(color *Color, style string) *FacialHair {
	return &FacialHair{Color: color, Style: style}
}
