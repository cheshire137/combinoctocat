package octocat

type Headgear struct {
	Style string
}

func NoHeadgear() *Headgear {
	return &Headgear{}
}

func NewHeadgear(style string) *Headgear {
	return &Headgear{Style: style}
}

func (h *Headgear) String() string {
	if len(h.Style) > 0 {
		return h.Style
	}
	return "no headgear"
}
