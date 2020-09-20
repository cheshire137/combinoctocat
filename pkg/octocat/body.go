package octocat

import "strings"

type Body struct {
	Colors []*Color
}

func NewBody(colors ...*Color) *Body {
	return &Body{Colors: colors}
}

func (b *Body) String() string {
	var sb strings.Builder

	sb.WriteString("Body: ")

	lastIndex := len(b.Colors) - 1
	for i, color := range b.Colors {
		sb.WriteString(color.String())
		if i < lastIndex {
			sb.WriteString(", ")
		}
	}

	return sb.String()
}
