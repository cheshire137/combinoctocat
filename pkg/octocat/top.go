package octocat

type Top struct {
	Style string
}

func NewTop(style string) *Top {
	return &Top{Style: style}
}
