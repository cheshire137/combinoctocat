package octocat

type Eyewear struct {
	Style string
}

func NewEyewear(style string) *Eyewear {
	return &Eyewear{Style: style}
}
