package octocat

import "strings"

type Octocat struct {
	Body         *Body
	Face         *Face
	Eyes         *Eyes
	Mouth        *Mouth
	Hair         *Hair
	FacialHair   *FacialHair
	Outfit       *Outfit
	AccessorySet *AccessorySet
	Prop         *Prop
}

func NewOctocat(body *Body, face *Face, eyes *Eyes, mouth *Mouth) *Octocat {
	return &Octocat{
		Body:  body,
		Face:  face,
		Eyes:  eyes,
		Mouth: mouth,
	}
}

func (o *Octocat) String() string {
	var sb strings.Builder
	sb.WriteString("Octocat: ")
	sb.WriteString(o.Body.String())
	sb.WriteString(o.Face.String())
	sb.WriteString(o.Eyes.String())
	sb.WriteString(o.Mouth.String())

	if o.Hair != nil {
		sb.WriteString(o.Hair.String())
	}

	if o.FacialHair != nil {
		sb.WriteString(o.FacialHair.String())
	}

	if o.Outfit != nil {
		sb.WriteString(o.Outfit.String())
	}

	if o.Prop != nil {
		sb.WriteString(o.Prop.String())
	}

	sb.WriteString(o.AccessorySet.String())

	return sb.String()
}
