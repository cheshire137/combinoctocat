package parse

import (
	"regexp"

	"github.com/cheshire137/combinoctocat/pkg/utils"
	"golang.org/x/net/html"
)

func GetElementById(n *html.Node, id string) *html.Node {
	return traverseById(n, id)
}

func GetElementsByClass(n *html.Node, className string) []*html.Node {
	return traverseByClass(n, className)
}

func GetAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func checkId(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		s, ok := GetAttribute(n, "id")
		if ok {
			return s == id
		}
	}
	return false
}

func checkClass(n *html.Node, className string) bool {
	if n.Type == html.ElementNode {
		s, ok := GetAttribute(n, "class")
		if ok {
			classNames := regexp.MustCompile(`\s+`).Split(s, -1)
			_, includesClass := utils.FindStr(classNames, className)
			return includesClass
		}
	}
	return false
}

func traverseByClass(n *html.Node, className string) []*html.Node {
	matches := []*html.Node{}

	if checkClass(n, className) {
		matches = append(matches, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		matches = append(matches, traverseByClass(c, className)...)
	}

	return matches
}

func traverseById(n *html.Node, id string) *html.Node {
	if checkId(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := traverseById(c, id)
		if result != nil {
			return result
		}
	}

	return nil
}
