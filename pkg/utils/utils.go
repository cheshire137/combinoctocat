package utils

import (
	"fmt"

	"github.com/cheshire137/combinoctocat/pkg/octocat"
)

func PrintColorList(colors []*octocat.Color) {
	lastIndex := len(colors) - 1

	for i, color := range colors {
		fmt.Print(color.String())
		if i < lastIndex {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}
}

func FindStr(haystack []string, needle string) (int, bool) {
	for i, item := range haystack {
		if item == needle {
			return i, true
		}
	}
	return -1, false
}
