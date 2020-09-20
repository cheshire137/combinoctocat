package octocat

type Top struct {
	Style string
}

func NewTop(style string) *Top {
	return &Top{Style: style}
}

func (t *Top) String() string {
	return t.Style
}
