package octocat

type Hair struct {
	Color *Color
	Style string
}

func NewHair(color *Color, style string) *Hair {
	return &Hair{Color: color, Style: style}
}
