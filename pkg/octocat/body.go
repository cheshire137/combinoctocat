package models

type Body struct {
	Colors []*Color
}

func NewBody(colors ...*Color) *Body {
	return &Body{Colors: colors}
}
