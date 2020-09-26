package octocat

type Bottom struct {
	Style string
}

func NoBottom() *Bottom {
	return &Bottom{}
}

func NewBottom(style string) *Bottom {
	return &Bottom{Style: style}
}

func (b *Bottom) String() string {
	if len(b.Style) > 0 {
		return b.Style
	}
	return "no bottom"
}
