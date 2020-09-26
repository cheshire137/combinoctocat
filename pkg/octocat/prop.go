package octocat

type Prop struct {
	Style string
}

func NewProp(style string) *Prop {
	return &Prop{Style: style}
}

func (p *Prop) String() string {
	return "Prop: " + p.Style
}
