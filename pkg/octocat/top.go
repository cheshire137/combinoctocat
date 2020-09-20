package octocat

type Top struct {
	Style string
}

func NoTop() *Top {
	return &Top{}
}

func NewTop(style string) *Top {
	return &Top{Style: style}
}

func (t *Top) String() string {
	if len(t.Style) > 0 {
		return "Top: " + t.Style
	}
	return "no top"
}
