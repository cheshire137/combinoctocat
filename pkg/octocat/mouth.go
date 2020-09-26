package octocat

type Mouth struct {
	Style string
}

func NewMouth(style string) *Mouth {
	return &Mouth{Style: style}
}

func (m *Mouth) String() string {
	return "Mouth: " + m.Style
}
