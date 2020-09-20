package octocat

type Bottom struct {
	Style string
}

func NewBottom(style string) *Bottom {
	return &Bottom{Style: style}
}

func (b *Bottom) String() string {
	return b.Style
}
