package octocat

import "strings"

type Eyes struct {
	Color *Color
	Style string
}

func NewEyes(color *Color, style string) *Eyes {
	return &Eyes{Color: color, Style: style}
}

func (e *Eyes) String() string {
	var sb strings.Builder
	sb.WriteString(e.Style)
	sb.WriteString(e.Color.String())
	return sb.String()
}
