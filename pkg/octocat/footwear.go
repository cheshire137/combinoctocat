package octocat

type Footwear struct {
	Style string
}

func NoFootwear() *Footwear {
	return &Footwear{}
}

func NewFootwear(style string) *Footwear {
	return &Footwear{Style: style}
}

func (f *Footwear) String() string {
	if len(f.Style) > 0 {
		return "Footwear: " + f.Style
	}
	return "no footwear"
}
