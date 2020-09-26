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

func NewOctocat(body *Body, face *Face, eyes *Eyes, mouth *Mouth, hair *Hair, facialHair *FacialHair, outfit *Outfit, accessorySet *AccessorySet, prop *Prop) *Octocat {
	return &Octocat{
		Body:         body,
		Face:         face,
		Eyes:         eyes,
		Mouth:        mouth,
		Hair:         hair,
		FacialHair:   facialHair,
		Outfit:       outfit,
		AccessorySet: accessorySet,
		Prop:         prop,
	}
}

func (o *Octocat) String() string {
	var sb strings.Builder
	sb.WriteString("Octocat: ")
	sb.WriteString(o.Body.String())
	sb.WriteString("\n")

	sb.WriteString("- ")
	sb.WriteString(o.Face.String())
	sb.WriteString("\n")

	sb.WriteString("- ")
	sb.WriteString(o.Eyes.String())
	sb.WriteString("\n")

	sb.WriteString("- ")
	sb.WriteString(o.Mouth.String())
	sb.WriteString("\n")

	if o.Hair != nil {
		sb.WriteString("- ")
		sb.WriteString(o.Hair.String())
		sb.WriteString("\n")
	}

	if o.FacialHair != nil {
		sb.WriteString("- ")
		sb.WriteString(o.FacialHair.String())
		sb.WriteString("\n")
	}

	if o.Outfit != nil {
		sb.WriteString("- ")
		sb.WriteString(o.Outfit.String())
		sb.WriteString("\n")
	}

	if o.Prop != nil {
		sb.WriteString("- ")
		sb.WriteString(o.Prop.String())
		sb.WriteString("\n")
	}

	sb.WriteString("- ")
	sb.WriteString(o.AccessorySet.String())

	return sb.String()
}
