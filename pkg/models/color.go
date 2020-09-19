package models

type Color struct {
	HexCode string
}

func NewColor(hexCode string) *Color {
	return &Color{HexCode: hexCode}
}
