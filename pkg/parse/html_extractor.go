package parse

import (
	"github.com/cheshire137/combinoctocat/pkg/octocat"
	"golang.org/x/net/html"
)

func ExtractBodies(rootNode *html.Node) []*octocat.Body {
	bodyNodes := GetElementsByClass(rootNode, "color-swatch")
	bodies := make([]*octocat.Body, len(bodyNodes))
	for i, node := range bodyNodes {
		colors := extractBodyColorsFromNode(node)
		bodies[i] = octocat.NewBody(colors...)
	}
	return bodies
}

func ExtractEyeColors(rootNode *html.Node) []*octocat.Color {
	return extractSingleColorFromChildNodes(rootNode)
}

func ExtractEyeStyles(rootNode *html.Node) []string {
	return extractImgAlts(rootNode)
}

func ExtractFaces(rootNode *html.Node) []*octocat.Face {
	faceNodes := GetElementsByClass(rootNode, "color-swatch")
	faces := make([]*octocat.Face, len(faceNodes))
	for i, node := range faceNodes {
		var faceColor *octocat.Color
		var noseColor *octocat.Color

		color1, ok := GetAttribute(node, "data-color")
		if ok {
			faceColor = octocat.NewColor(color1)
		} else {
			faceColor = nil
		}

		color2, ok := GetAttribute(node, "data-color-nose")
		if ok {
			noseColor = octocat.NewColor(color2)
		} else {
			noseColor = nil
		}

		faces[i] = octocat.NewFace(faceColor, noseColor)
	}
	return faces
}

func ExtractHairColors(rootNode *html.Node) []*octocat.Color {
	return extractSingleColorFromChildNodes(rootNode)
}

func ExtractFacialHairColors(rootNode *html.Node) []*octocat.Color {
	return extractSingleColorFromChildNodes(rootNode)
}

func ExtractHairStyles(rootNode *html.Node) []string {
	return extractImgAlts(rootNode)
}

func ExtractHeadgears(rootNode *html.Node) []*octocat.Headgear {
	styles := extractImgAlts(rootNode)
	headgears := make([]*octocat.Headgear, len(styles))
	for i, style := range styles {
		headgears[i] = octocat.NewHeadgear(style)
	}
	return headgears
}

func extractImgAlts(rootNode *html.Node) []string {
	previewNodes := GetElementsByClass(rootNode, "preview")
	imgAltAttributes := make([]string, len(previewNodes))
	for i, node := range previewNodes {
		altAttribute, ok := GetAttribute(node, "alt")
		if ok {
			imgAltAttributes[i] = altAttribute
		}
	}
	return imgAltAttributes
}

func extractSingleColorFromChildNodes(rootNode *html.Node) []*octocat.Color {
	colorNodes := GetElementsByClass(rootNode, "color-swatch")
	colors := make([]*octocat.Color, len(colorNodes))
	for i, node := range colorNodes {
		color, ok := GetAttribute(node, "data-color")
		if ok {
			colors[i] = octocat.NewColor(color)
		}
	}
	return colors
}

func extractBodyColorsFromNode(node *html.Node) []*octocat.Color {
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
