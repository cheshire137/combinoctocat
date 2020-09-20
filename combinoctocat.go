package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cheshire137/combinoctocat/pkg/octocat"
	"github.com/cheshire137/combinoctocat/pkg/options"
	"github.com/cheshire137/combinoctocat/pkg/parse"
	"github.com/cheshire137/combinoctocat/pkg/utils"
	"golang.org/x/net/html"
)

func main() {
	options, err := options.ParseOptions()
	if err != nil {
		os.Exit(1)
		return
	}

	f, err := os.Open(options.InputPath)
	if err != nil {
		fmt.Println("Could not read file: " + err.Error())
		os.Exit(1)
		return
	}

	reader := bufio.NewReader(f)
	node, err := html.Parse(reader)
	if err != nil {
		fmt.Println("Could not parse HTML: " + err.Error())
		os.Exit(1)
		return
	}

	bodyColorNode := parse.GetElementById(node, "body-color")
	bodyChoices := parse.ExtractBodies(bodyColorNode)

	eyeStyleNode := parse.GetElementById(node, "cp-eyes")
	eyeStyleChoices := parse.ExtractEyeStyles(eyeStyleNode)

	eyeColorNode := parse.GetElementById(node, "eye-color")
	eyeColorChoices := parse.ExtractEyeColors(eyeColorNode)

	faceNode := parse.GetElementById(node, "face-color")
	faceChoices := parse.ExtractFaces(faceNode)

	hairStyleNode := parse.GetElementById(node, "cp-hair")
	hairStyleChoices := parse.ExtractHairStyles(hairStyleNode)

	hairColorNode := parse.GetElementById(node, "hair-color")
	hairColorChoices := parse.ExtractHairColors(hairColorNode)

	facialHairColorNode := parse.GetElementById(node, "facehair-color")
	facialHairColorChoices := parse.ExtractFacialHairColors(facialHairColorNode)

	headgearNode := parse.GetElementById(node, "cp-headgear")
	headgearChoices := parse.ExtractHeadgears(headgearNode)

	hairChoices := []*octocat.Hair{}
	for _, style := range hairStyleChoices {
		for _, color := range hairColorChoices {
			hairChoices = append(hairChoices, octocat.NewHair(color, style))
		}
	}

	eyeChoices := []*octocat.Eyes{}
	for _, style := range eyeStyleChoices {
		for _, color := range eyeColorChoices {
			eyeChoices = append(eyeChoices, octocat.NewEyes(color, style))
		}
	}

	fmt.Printf("Octocat body choices (%d):\n", len(bodyChoices))
	for _, body := range bodyChoices {
		fmt.Println("- " + body.String())
	}

	fmt.Printf("\n%d Octocat eye choices\n", len(eyeChoices))

	fmt.Printf("\nOctocat face choices (%d):\n", len(faceChoices))
	for _, face := range faceChoices {
		fmt.Println("- " + face.String())
	}

	fmt.Printf("\n%d Octocat hair choices\n", len(hairChoices))

	fmt.Printf("\nOctocat facial hair color choices (%d):\n", len(facialHairColorChoices))
	utils.PrintColorList(facialHairColorChoices)

	fmt.Printf("\nOctocat headgear choices (%d):\n", len(headgearChoices))
	utils.PrintHeadgearList(headgearChoices)

	// topNode := parse.GetElementById(node, "cp-tops")

	// bottomNode := parse.GetElementById(node, "cp-bottoms")

	// footwearNode := parse.GetElementById(node, "cp-footwear")

	// eyewearNode := parse.GetElementById(node, "cp-eyewear")

	// propNode := parse.GetElementById(node, "cp-props")

	// facialHairStyleNode := parse.GetElementById(node, "cp-faceHair")
}
