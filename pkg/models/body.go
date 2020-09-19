package models

type Body struct {
	Color *Color
}

func NewBody(color *Color) *Body {
	return &Body{Color: color}
}
