package octocat

type Eyewear struct {
	Style string
}

func NoEyewear() *Eyewear {
	return &Eyewear{}
}

func NewEyewear(style string) *Eyewear {
	return &Eyewear{Style: style}
}

func (e *Eyewear) String() string {
	if len(e.Style) > 0 {
		return e.Style
	}
	return "no eyewear"
}
