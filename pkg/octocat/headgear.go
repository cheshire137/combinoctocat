package octocat

type Headgear struct {
	Style string
}

func NewHeadgear(style string) *Headgear {
	return &Headgear{Style: style}
}

func (h *Headgear) String() string {
	return h.Style
}
