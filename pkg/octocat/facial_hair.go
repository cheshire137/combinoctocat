package octocat

import "strings"

type FacialHair struct {
	Color *Color
	Style string
}

func NewFacialHair(color *Color, style string) *FacialHair {
	return &FacialHair{Color: color, Style: style}
}

func (h *FacialHair) String() string {
	var sb strings.Builder
	sb.WriteString("Facial hair: ")
	sb.WriteString(h.Style)
	if h.Color != nil {
		sb.WriteString(" / ")
		sb.WriteString(h.Color.String())
	}
	return sb.String()
}
