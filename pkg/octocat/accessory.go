package octocat

type Accessory struct {
	Style string
}

func NewAccessory(style string) *Accessory {
	return &Accessory{Style: style}
}

func (a *Accessory) String() string {
	return a.Style
}
