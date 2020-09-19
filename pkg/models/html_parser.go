package models

import (
	"bufio"

	"golang.org/x/net/html"
)

type HtmlParser struct {
	Node *html.Node
}

func ParseHtml(reader *bufio.Reader) (*HtmlParser, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	return &HtmlParser{Node: doc}, nil
}
