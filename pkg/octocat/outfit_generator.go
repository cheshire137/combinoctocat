package octocat

import "math/rand"

type OutfitGenerator struct {
	Tops            []*Top
	Bottoms         []*Bottom
	FootwearChoices []*Footwear
	HeadgearChoices []*Headgear
	EyewearChoices  []*Eyewear
}

func NewOutfitGenerator(t []*Top, b []*Bottom, f []*Footwear, h []*Headgear, e []*Eyewear) *OutfitGenerator {
	return &OutfitGenerator{
		Tops:            t,
		Bottoms:         b,
		FootwearChoices: f,
		HeadgearChoices: h,
		EyewearChoices:  e,
	}
}

func (g *OutfitGenerator) Generate() *Outfit {
	top := g.randomTop()
	bottom := g.randomBottom()
	footwear := g.randomFootwear()
	headgear := g.randomHeadgear()
	eyewear := g.randomEyewear()
	return NewOutfit(top, bottom, footwear, headgear, eyewear)
}

func (g *OutfitGenerator) TotalPossible() uint64 {
	return uint64(len(g.Tops) * len(g.Bottoms) * len(g.FootwearChoices) * len(g.HeadgearChoices) * len(g.EyewearChoices))
}

func (g *OutfitGenerator) randomTop() *Top {
	return g.Tops[rand.Intn(len(g.Tops))]
}

func (g *OutfitGenerator) randomBottom() *Bottom {
	return g.Bottoms[rand.Intn(len(g.Bottoms))]
}

func (g *OutfitGenerator) randomFootwear() *Footwear {
	return g.FootwearChoices[rand.Intn(len(g.FootwearChoices))]
}

func (g *OutfitGenerator) randomHeadgear() *Headgear {
	return g.HeadgearChoices[rand.Intn(len(g.HeadgearChoices))]
}

func (g *OutfitGenerator) randomEyewear() *Eyewear {
	return g.EyewearChoices[rand.Intn(len(g.EyewearChoices))]
}
