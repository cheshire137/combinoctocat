package parse

import (
	"fmt"

	"github.com/cheshire137/combinoctocat/pkg/octocat"
	"golang.org/x/net/html"
)

func ExtractBodies(node *html.Node) []*octocat.Body {
	bodyNodes := GetElementsByClass(node, "color-swatch")
	fmt.Printf("got %d body colors\n", len(bodyNodes))
	bodies := make([]*octocat.Body, len(bodyNodes))
	return bodies
}
