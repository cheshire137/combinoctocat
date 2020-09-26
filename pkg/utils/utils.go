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

func PrintEyewearList(eyewears []*octocat.Eyewear) {
	strList := make([]string, len(eyewears))
	for i, eyewear := range eyewears {
		strList[i] = eyewear.String()
	}
	PrintCommaSeparatedList(strList)
}

func PrintMouthList(mouths []*octocat.Mouth) {
	strList := make([]string, len(mouths))
	for i, mouth := range mouths {
		strList[i] = mouth.String()
	}
	PrintCommaSeparatedList(strList)
}

func PrintPropList(props []*octocat.Prop) {
	strList := make([]string, len(props))
	for i, prop := range props {
		strList[i] = prop.String()
	}
	PrintCommaSeparatedList(strList)
}

func PrintAccessoryList(accessories []*octocat.Accessory) {
	strList := make([]string, len(accessories))
	for i, accessory := range accessories {
		strList[i] = accessory.String()
	}
	PrintCommaSeparatedList(strList)
}

func PrintFootwearList(footwears []*octocat.Footwear) {
	strList := make([]string, len(footwears))
	for i, footwear := range footwears {
		strList[i] = footwear.String()
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

func PrintBottomList(bottoms []*octocat.Bottom) {
	strList := make([]string, len(bottoms))
	for i, bottom := range bottoms {
		strList[i] = bottom.String()
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

func GetPositiveEnglishNumberName(number uint64) string {
	englishNames := map[uint64]string{
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		15: "fifteen",
		18: "eighteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}

	var thousand uint64
	var million uint64
	var billion uint64
	var trillion uint64
	var quadrillion uint64
	var quintrillion uint64
	var sextillion uint64

	thousand = 1000
	million = thousand * thousand
	billion = million * thousand
	trillion = billion * thousand
	quadrillion = trillion * thousand
	quintrillion = quadrillion * thousand
	sextillion = quintrillion * thousand

	name, ok := englishNames[number]
	if ok {
		return name
	}

	if number < 20 {
		return englishNames[number-10] + "teen"
	}

	if number < 100 {
		if number%10 == 0 {
			return englishNames[number]
		}
		name, ok := englishNames[number/10*10]
		if ok {
			return name + "-" + englishNames[number%10]
		}
	}

	if number < thousand {
		if number%100 == 0 {
			return englishNames[number] + " hundred"
		}
		name, ok := englishNames[number/100]
		if ok {
			return name + " hundred " + GetPositiveEnglishNumberName(number%100)
		}
	}

	if number < million {
		prefix := GetPositiveEnglishNumberName(number/thousand) + " thousand"
		if number%thousand == 0 {
			return prefix
		}
		return prefix + " " + GetPositiveEnglishNumberName(number%thousand)
	}

	if number < billion {
		prefix := GetPositiveEnglishNumberName(number/million) + " million"
		if number%million == 0 {
			return prefix
		}
		return prefix + " " + GetPositiveEnglishNumberName(number%million)
	}

	if number < trillion {
		prefix := GetPositiveEnglishNumberName(number/billion) + " billion"
		if number%billion == 0 {
			return prefix
		}
		return prefix + " " + GetPositiveEnglishNumberName(number%billion)
	}

	if number < quadrillion {
		prefix := GetPositiveEnglishNumberName(number/trillion) + " trillion"
		if number%trillion == 0 {
			return prefix
		}
		return prefix + " " + GetPositiveEnglishNumberName(number%trillion)
	}

	if number < quintrillion {
		prefix := GetPositiveEnglishNumberName(number/quadrillion) + " quadrillion"
		if number%quadrillion == 0 {
			return prefix
		}
		return prefix + " " + GetPositiveEnglishNumberName(number%quadrillion)
	}

	if number < sextillion {
		prefix := GetPositiveEnglishNumberName(number/quintrillion) + " quintrillion"
		if number%quintrillion == 0 {
			return prefix
		}
		return prefix + " " + GetPositiveEnglishNumberName(number%quintrillion)
	}

	return fmt.Sprintf("number is too large: %d", number)
}
