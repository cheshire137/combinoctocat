package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cheshire137/combinoctocat/pkg/options"
	"github.com/cheshire137/combinoctocat/pkg/parse"
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
	totalBodyChoices := len(bodyChoices)

	// eyeStyleNode := parse.GetElementById(node, "cp-eyes")
	eyeColorNode := parse.GetElementById(node, "eye-color")
	eyeColorChoices := parse.ExtractEyeColors(eyeColorNode)
	totalEyeColorChoices := len(eyeColorChoices)

	faceNode := parse.GetElementById(node, "face-color")
	faceChoices := parse.ExtractFaces(faceNode)
	totalFaceChoices := len(faceChoices)

	fmt.Printf("Octocat body choices (%d):\n", totalBodyChoices)
	for _, body := range bodyChoices {
		fmt.Println("- " + body.String())
	}

	fmt.Printf("\nOctocat eye color choices (%d):\n", totalEyeColorChoices)
	lastEyeColorIndex := totalEyeColorChoices - 1
	for i, color := range eyeColorChoices {
		fmt.Print(color.String())
		if i < lastEyeColorIndex {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}

	fmt.Printf("\nOctocat face choices (%d):\n", totalFaceChoices)
	for _, face := range faceChoices {
		fmt.Println("- " + face.String())
	}

	// hairStyleNode := parse.GetElementById(node, "cp-hair")
	// hairColorNode := parse.GetElementById(node, "hair-color")

	// headgearNode := parse.GetElementById(node, "cp-headgear")

	// topNode := parse.GetElementById(node, "cp-tops")

	// bottomNode := parse.GetElementById(node, "cp-bottoms")

	// footwearNode := parse.GetElementById(node, "cp-footwear")

	// eyewearNode := parse.GetElementById(node, "cp-eyewear")

	// propNode := parse.GetElementById(node, "cp-props")

	// facialHairStyleNode := parse.GetElementById(node, "cp-faceHair")
	// facialHairColorNode := parse.GetElementById(node, "facehair-color")
}
