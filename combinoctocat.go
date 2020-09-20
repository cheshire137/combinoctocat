package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/cheshire137/combinoctocat/pkg/octocat"
	"github.com/cheshire137/combinoctocat/pkg/options"
	"github.com/cheshire137/combinoctocat/pkg/parse"
	"github.com/cheshire137/combinoctocat/pkg/utils"
	"golang.org/x/net/html"
	"golang.org/x/text/message"
)

func main() {
	rand.Seed(time.Now().UnixNano())

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

	facialHairStyleNode := parse.GetElementById(node, "cp-faceHair")
	facialHairStyleChoices := parse.ExtractFacialHairStyles(facialHairStyleNode)

	headgearNode := parse.GetElementById(node, "cp-headgear")
	headgearChoices := parse.ExtractHeadgears(headgearNode)

	mouthNode := parse.GetElementById(node, "cp-mouths")
	mouthChoices := parse.ExtractMouths(mouthNode)

	topNode := parse.GetElementById(node, "cp-tops")
	topChoices := parse.ExtractTops(topNode)

	bottomNode := parse.GetElementById(node, "cp-bottoms")
	bottomChoices := parse.ExtractBottoms(bottomNode)

	footwearNode := parse.GetElementById(node, "cp-footwear")
	footwearChoices := parse.ExtractFootwears(footwearNode)

	eyewearNode := parse.GetElementById(node, "cp-eyewear")
	eyewearChoices := parse.ExtractEyewears(eyewearNode)

	propNode := parse.GetElementById(node, "cp-props")
	propChoices := parse.ExtractProps(propNode)

	accessoriesNode := parse.GetElementById(node, "cp-misc")
	accessoryChoices := parse.ExtractAccessories(accessoriesNode)

	hairChoices := []*octocat.Hair{octocat.NoHair()}
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

	facialHairChoices := []*octocat.FacialHair{octocat.NoFacialHair()}
	for _, style := range facialHairStyleChoices {
		for _, color := range facialHairColorChoices {
			facialHairChoices = append(facialHairChoices, octocat.NewFacialHair(color, style))
		}
	}

	totalPossibleOutfits := len(topChoices) * len(bottomChoices) * len(footwearChoices) * len(headgearChoices) * len(eyewearChoices)

	// octocats := []*octocat.Octocat{}
	// for _, body := range bodyChoices {
	// 	for _, prop := range propChoices {
	// 		for _, eyes := range eyeChoices {
	// 			for _, facialHair := range facialHairChoices {
	// 				for _, hair := range hairChoices {
	// 					for _, mouth := range mouthChoices {
	// 						for _, face := range faceChoices {
	// 							octocat := octocat.NewOctocat(body, face, eyes, mouth)
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

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

	fmt.Printf("\n%d Octocat facial hair choices\n", len(facialHairChoices))

	fmt.Printf("\nOctocat mouth choices (%d):\n", len(mouthChoices))
	utils.PrintMouthList(mouthChoices)

	fmt.Printf("\nOctocat prop choices (%d):\n", len(propChoices))
	utils.PrintPropList(propChoices)

	fmt.Printf("\nOctocat accessory choices (%d):\n", len(accessoryChoices))
	utils.PrintAccessoryList(accessoryChoices)

	fmt.Print("\n")
	printer := message.NewPrinter(message.MatchLanguage("en"))
	printer.Print(totalPossibleOutfits)
	fmt.Printf(" outfit choices (%d tops, %d bottoms, %d headgears, %d eyewears, %d footgears)\n",
		len(topChoices), len(bottomChoices), len(headgearChoices), len(eyewearChoices), len(footwearChoices))
}
