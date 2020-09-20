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
	if h.Color != nil {
		sb.WriteString(h.Color.String())
		sb.WriteString(" / ")
	}
	sb.WriteString(h.Style)
	return sb.String()
}
