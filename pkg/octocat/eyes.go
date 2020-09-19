package octocat

type Eyes struct {
	Color *Color
	Style string
}

func NewEyes(color *Color, style string) *Eyes {
	return &Eyes{Color: color, Style: style}
}
