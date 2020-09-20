package octocat

import "strings"

type FacialHair struct {
	Color *Color
	Style string
}

func NoFacialHair() *FacialHair {
	return &FacialHair{}
}

func NewFacialHair(color *Color, style string) *FacialHair {
	return &FacialHair{Color: color, Style: style}
}

func (h *FacialHair) String() string {
	if h.Color == nil || len(h.Style) < 1 {
		return "no facial hair"
	}

	var sb strings.Builder
	sb.WriteString("Facial hair: ")
	sb.WriteString(h.Style)
	sb.WriteString(" / ")
	sb.WriteString(h.Color.String())
	return sb.String()
}
