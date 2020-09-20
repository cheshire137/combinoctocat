package parse

import (
	"github.com/cheshire137/combinoctocat/pkg/octocat"
	"golang.org/x/net/html"
)

func ExtractBodies(node *html.Node) []*octocat.Body {
	bodyNodes := GetElementsByClass(node, "color-swatch")
	bodies := make([]*octocat.Body, len(bodyNodes))
	for i, node := range bodyNodes {
		colors := ExtractColorsFromNode(node)
		bodies[i] = octocat.NewBody(colors...)
	}
	return bodies
}

func ExtractColorsFromNode(node *html.Node) []*octocat.Color {
	colors := []*octocat.Color{}

	color1, ok := GetAttribute(node, "data-color")
	if ok {
		colors = append(colors, octocat.NewColor(color1))
	}

	color2, ok := GetAttribute(node, "data-second-color")
	if ok {
		colors = append(colors, octocat.NewColor(color2))
	}

	color3, ok := GetAttribute(node, "data-third-color")
	if ok {
		colors = append(colors, octocat.NewColor(color3))
	}

	return colors
}
