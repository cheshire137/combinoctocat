package parse

import (
	"strings"

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

func ExtractEyewears(rootNode *html.Node) []*octocat.Eyewear {
	styles := extractImgAlts(rootNode)
	eyeWears := make([]*octocat.Eyewear, len(styles))
	for i, style := range styles {
		eyeWears[i] = octocat.NewEyewear(style)
	}
	return eyeWears
}

func ExtractTops(rootNode *html.Node) []*octocat.Top {
	imgAlts := extractImgAlts(rootNode)
	imgSrcs := extractImgSrcs(rootNode)
	tops := make([]*octocat.Top, len(imgAlts))
	for i, imgAlt := range imgAlts {
		style := getNameFromImgAltAndSrc(imgAlt, imgSrcs[i], "tops-")
		tops[i] = octocat.NewTop(style)
	}
	return tops
}

func ExtractBottoms(rootNode *html.Node) []*octocat.Bottom {
	imgAlts := extractImgAlts(rootNode)
	bottoms := make([]*octocat.Bottom, len(imgAlts))
	for i, imgAlt := range imgAlts {
		bottoms[i] = octocat.NewBottom(imgAlt)
	}
	return bottoms
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

func ExtractFacialHairStyles(rootNode *html.Node) []string {
	return extractImgAlts(rootNode)
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

func ExtractFootwears(rootNode *html.Node) []*octocat.Footwear {
	imgAlts := extractImgAlts(rootNode)
	imgSrcs := extractImgSrcs(rootNode)
	footwears := make([]*octocat.Footwear, len(imgAlts))
	for i, imgAlt := range imgAlts {
		style := getNameFromImgAltAndSrc(imgAlt, imgSrcs[i], "footwear-")
		footwears[i] = octocat.NewFootwear(style)
	}
	return footwears
}

func getNameFromImgAltAndSrc(imgAlt string, imgSrc string, prefix string) string {
	if len(imgAlt) > 0 {
		return imgAlt
	}

	srcParts := strings.Split(imgSrc, "/")
	totalSrcParts := len(srcParts)
	if totalSrcParts > 0 {
		fileName := srcParts[totalSrcParts-1]
		fileNameWithoutExtension := strings.TrimSuffix(fileName, ".svg")
		return strings.TrimPrefix(fileNameWithoutExtension, prefix)
	}

	return imgSrc
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

func extractImgSrcs(rootNode *html.Node) []string {
	previewNodes := GetElementsByClass(rootNode, "preview")
	imgSrcAttributes := make([]string, len(previewNodes))
	for i, node := range previewNodes {
		srcAttribute, ok := GetAttribute(node, "src")
		if ok {
			imgSrcAttributes[i] = srcAttribute
		}
	}
	return imgSrcAttributes
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
