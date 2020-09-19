package models

type Prop struct {
	Style string
}

func NewProp(style string) *Prop {
	return &Prop{Style: style}
}
