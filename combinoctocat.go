package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cheshire137/combinoctocat/pkg/models"
	"golang.org/x/net/html"
)

func main() {
	options := models.ParseOptions()
	if len(options.InputPath) < 1 {
		models.DisplayOptions()
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

	bodyColorNode := models.GetElementById(node, "body-color")

	eyeStyleNode := models.GetElementById(node, "cp-eyes")
	eyeColorNode := models.GetElementById(node, "eye-color")

	faceColorNode := models.GetElementById(node, "face-color")

	hairStyleNode := models.GetElementById(node, "cp-hair")
	hairColorNode := models.GetElementById(node, "hair-color")

	headgearNode := models.GetElementById(node, "cp-headgear")

	topNode := models.GetElementById(node, "cp-tops")

	bottomNode := models.GetElementById(node, "cp-bottoms")

	footwearNode := models.GetElementById(node, "cp-footwear")

	eyewearNode := models.GetElementById(node, "cp-eyewear")

	propNode := models.GetElementById(node, "cp-props")

	facialHairStyleNode := models.GetElementById(node, "cp-faceHair")
	facialHairColorNode := models.GetElementById(node, "facehair-color")
}
