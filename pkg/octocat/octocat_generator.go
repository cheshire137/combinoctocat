package octocat

import "math/rand"

type OctocatGenerator struct {
	Bodies          []*Body
	Faces           []*Face
	EyeChoices      []*Eyes
	Mouths          []*Mouth
	Hairs           []*Hair
	FacialHairs     []*FacialHair
	AccessorySets   []*AccessorySet
	Props           []*Prop
	OutfitGenerator *OutfitGenerator
}

func NewOctocatGenerator(b []*Body, f []*Face, e []*Eyes, m []*Mouth, h []*Hair, fh []*FacialHair, a []*AccessorySet, p []*Prop, g *OutfitGenerator) *OctocatGenerator {
	return &OctocatGenerator{
		Bodies:          b,
		Faces:           f,
		EyeChoices:      e,
		Mouths:          m,
		Hairs:           h,
		FacialHairs:     fh,
		AccessorySets:   a,
		Props:           p,
		OutfitGenerator: g,
	}
}

func (g *OctocatGenerator) Generate() *Octocat {
	outfit := g.OutfitGenerator.Generate()
	body := g.randomBody()
	face := g.randomFace()
	eyes := g.randomEyes()
	mouth := g.randomMouth()
	hair := g.randomHair()
	facialHair := g.randomFacialHair()
	accessorySet := g.randomAccessorySet()
	prop := g.randomProp()
	return NewOctocat(body, face, eyes, mouth, hair, facialHair, outfit, accessorySet, prop)
}

func (g *OctocatGenerator) randomBody() *Body {
	return g.Bodies[rand.Intn(len(g.Bodies))]
}

func (g *OctocatGenerator) randomFace() *Face {
	return g.Faces[rand.Intn(len(g.Faces))]
}

func (g *OctocatGenerator) randomEyes() *Eyes {
	return g.EyeChoices[rand.Intn(len(g.EyeChoices))]
}

func (g *OctocatGenerator) randomMouth() *Mouth {
	return g.Mouths[rand.Intn(len(g.Mouths))]
}

func (g *OctocatGenerator) randomHair() *Hair {
	return g.Hairs[rand.Intn(len(g.Hairs))]
}

func (g *OctocatGenerator) randomFacialHair() *FacialHair {
	return g.FacialHairs[rand.Intn(len(g.FacialHairs))]
}

func (g *OctocatGenerator) randomAccessorySet() *AccessorySet {
	return g.AccessorySets[rand.Intn(len(g.AccessorySets))]
}

func (g *OctocatGenerator) randomProp() *Prop {
	return g.Props[rand.Intn(len(g.Props))]
}
