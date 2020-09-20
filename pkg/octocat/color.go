package octocat

type Color struct {
	HexCode string
}

func NewColor(hexCode string) *Color {
	return &Color{HexCode: hexCode}
}

func (c *Color) String() string {
	return c.HexCode
}
