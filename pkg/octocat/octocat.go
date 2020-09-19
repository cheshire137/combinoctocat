package models

type Octocat struct {
	Body        *Body
	Face        *Face
	Eyes        *Eyes
	Mouth       *Mouth
	Hair        *Hair
	FacialHair  *FacialHair
	Top         *Top
	Bottom      *Bottom
	Footwear    *Footwear
	Headgear    *Headgear
	Eyewear     *Eyewear
	Accessories []*Accessory
	Prop        *Prop
}
