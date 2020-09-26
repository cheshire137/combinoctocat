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

	accessorySets := octocat.GetAccessorySets(accessoryChoices)

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

	totalPossibleOutfits := uint64(len(topChoices) * len(bottomChoices) * len(footwearChoices) * len(headgearChoices) * len(eyewearChoices))

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

	fmt.Println("Octocat customization options:")

	totalBodies := uint64(len(bodyChoices))
	totalEyes := uint64(len(eyeChoices))
	totalFaces := uint64(len(faceChoices))
	totalHairs := uint64(len(hairChoices))
	totalFacialHairs := uint64(len(facialHairChoices))
	totalMouths := uint64(len(mouthChoices))
	totalProps := uint64(len(propChoices))
	totalAccessoryCombos := uint64(len(accessorySets))

	printer := message.NewPrinter(message.MatchLanguage("en"))

	fmt.Print("\n")
	fmt.Printf("%d body choices\n", totalBodies)
	fmt.Printf("%d eye choices\n", totalEyes)
	fmt.Printf("%d face choices\n", totalFaces)
	fmt.Printf("%d mouth choices\n", totalMouths)
	fmt.Printf("%d prop choices\n", totalProps)
	printer.Print(totalAccessoryCombos)
	fmt.Println(" accessory combinations")

	fmt.Printf("\n%d hair choices (including none)\n", totalHairs)
	fmt.Printf("- %d colors\n", len(hairColorChoices))
	fmt.Printf("- %d styles\n", len(hairStyleChoices))

	fmt.Printf("\n%d facial hair choices (including none)\n", totalFacialHairs)
	fmt.Printf("- %d colors\n", len(facialHairColorChoices))
	fmt.Printf("- %d styles\n", len(facialHairStyleChoices))

	fmt.Print("\n")
	printer.Print(totalPossibleOutfits)
	fmt.Println(" outfit choices")
	fmt.Printf("- %d tops\n", len(topChoices))
	fmt.Printf("- %d bottoms\n", len(bottomChoices))
	fmt.Printf("- %d headgears\n", len(headgearChoices))
	fmt.Printf("- %d eyewears\n", len(eyewearChoices))
	fmt.Printf("- %d footgears\n", len(footwearChoices))

	var totalOctocats uint64
	totalOctocats = totalBodies * totalEyes * totalFaces * totalHairs * totalFacialHairs * totalMouths * totalProps * totalAccessoryCombos * totalPossibleOutfits
	fmt.Print("\n")
	printer.Print(totalOctocats)
	fmt.Println(" possible Octocats")

	fmt.Println("\n" + utils.GetPositiveEnglishNumberName(totalOctocats))
}
