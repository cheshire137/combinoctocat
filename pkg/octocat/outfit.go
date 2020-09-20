package octocat

import (
	"math/rand"
	"strings"
)

type Outfit struct {
	Top      *Top
	Bottom   *Bottom
	Footwear *Footwear
	Headgear *Headgear
	Eyewear  *Eyewear
}

func NewOutfit(top *Top, bottom *Bottom, footwear *Footwear, headgear *Headgear, eyewear *Eyewear) *Outfit {
	return &Outfit{
		Top:      top,
		Bottom:   bottom,
		Footwear: footwear,
		Headgear: headgear,
		Eyewear:  eyewear,
	}
}

func (o *Outfit) String() string {
	parts := []string{}
	if o.Top != nil {
		parts = append(parts, o.Top.String())
	}
	if o.Bottom != nil {
		parts = append(parts, o.Bottom.String())
	}
	if o.Footwear != nil {
		parts = append(parts, o.Footwear.String())
	}
	if o.Headgear != nil {
		parts = append(parts, o.Headgear.String())
	}
	if o.Eyewear != nil {
		parts = append(parts, o.Eyewear.String())
	}
	return strings.Join(parts, " / ")
}

func GenerateOutfit(topChoices []*Top, bottomChoices []*Bottom, footwearChoices []*Footwear, headgearChoices []*Headgear, eyewearChoices []*Eyewear) *Outfit {
	topIndex := rand.Intn(len(topChoices))
	bottomIndex := rand.Intn(len(bottomChoices))
	footwearIndex := rand.Intn(len(footwearChoices))
	headgearIndex := rand.Intn(len(headgearChoices))
	eyewearIndex := rand.Intn(len(eyewearChoices))
	return NewOutfit(topChoices[topIndex], bottomChoices[bottomIndex], footwearChoices[footwearIndex],
		headgearChoices[headgearIndex], eyewearChoices[eyewearIndex])
}
