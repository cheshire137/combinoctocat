package octocat

import "strings"

type Hair struct {
	Color *Color
	Style string
}

func NewHair(color *Color, style string) *Hair {
	return &Hair{Color: color, Style: style}
}

func (h *Hair) String() string {
	var sb strings.Builder
	sb.WriteString("Hair: ")
	sb.WriteString(h.Style)
	if h.Color != nil {
		sb.WriteString(" / ")
		sb.WriteString(h.Color.String())
	}
	return sb.String()
}
