package octocat

type Mouth struct {
	Style string
}

func NewMouth(style string) *Mouth {
	return &Mouth{Style: style}
}
