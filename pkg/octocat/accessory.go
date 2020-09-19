package octocat

type Accessory struct {
	Style string
}

func NewAccessory(style string) *Accessory {
	return &Accessory{Style: style}
}
