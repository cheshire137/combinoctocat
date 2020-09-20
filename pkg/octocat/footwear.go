package octocat

type Footwear struct {
	Style string
}

func NewFootwear(style string) *Footwear {
	return &Footwear{Style: style}
}

func (f *Footwear) String() string {
	return f.Style
}
