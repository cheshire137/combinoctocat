package utils

import (
	"fmt"

	"github.com/cheshire137/combinoctocat/pkg/octocat"
)

func PrintCommaSeparatedList(list []string) {
	lastIndex := len(list) - 1

	for i, value := range list {
		fmt.Print(value)
		if i < lastIndex {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}
}

func PrintColorList(colors []*octocat.Color) {
	strList := make([]string, len(colors))
	for i, color := range colors {
		strList[i] = color.String()
	}
	PrintCommaSeparatedList(strList)
}

func PrintHeadgearList(headgears []*octocat.Headgear) {
	strList := make([]string, len(headgears))
	for i, headgear := range headgears {
		strList[i] = headgear.String()
	}
	PrintCommaSeparatedList(strList)
}

func PrintTopList(tops []*octocat.Top) {
	strList := make([]string, len(tops))
	for i, top := range tops {
		strList[i] = top.String()
	}
	PrintCommaSeparatedList(strList)
}

func FindStr(haystack []string, needle string) (int, bool) {
	for i, item := range haystack {
		if item == needle {
			return i, true
		}
	}
	return -1, false
}
