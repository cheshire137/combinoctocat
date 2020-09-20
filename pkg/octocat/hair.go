package octocat

import "strings"

type Hair struct {
	Color *Color
	Style string
}

func NoHair() *Hair {
	return &Hair{}
}

func NewHair(color *Color, style string) *Hair {
	return &Hair{Color: color, Style: style}
}

func (h *Hair) String() string {
	if h.Color == nil || len(h.Style) < 1 {
		return "no hair"
	}

	var sb strings.Builder
	sb.WriteString("Hair: ")
	sb.WriteString(h.Style)
	sb.WriteString(" / ")
	sb.WriteString(h.Color.String())
	return sb.String()
}
