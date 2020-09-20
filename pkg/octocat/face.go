package octocat

import "strings"

type Face struct {
	Color     *Color
	NoseColor *Color
}

func NewFace(color *Color, noseColor *Color) *Face {
	return &Face{Color: color, NoseColor: noseColor}
}

func (f *Face) String() string {
	var sb strings.Builder

	sb.WriteString("Face: ")

	if f.Color != nil {
		sb.WriteString(f.Color.String())
	}

	if f.NoseColor != nil {
		if f.Color != nil {
			sb.WriteString(" / ")
		}
		sb.WriteString("Nose: ")
		sb.WriteString(f.NoseColor.String())
	}

	return sb.String()
}
