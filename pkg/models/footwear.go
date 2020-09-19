package models

type Footwear struct {
	Style string
}

func NewFootwear(style string) *Footwear {
	return &Footwear{Style: style}
}
