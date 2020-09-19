package models

type Face struct {
	Color *Color
}

func NewFace(color *Color) *Face {
	return &Face{Color: color}
}
