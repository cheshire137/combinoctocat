package octocat

import "strings"

type AccessorySet struct {
	Accessories []*Accessory
}

func NewAccessorySet(accessories ...*Accessory) *AccessorySet {
	return &AccessorySet{Accessories: accessories}
}

func (s *AccessorySet) String() string {
	if len(s.Accessories) < 1 {
		return "no accessories"
	}
	strList := make([]string, len(s.Accessories))
	for i, accessory := range s.Accessories {
		strList[i] = accessory.String()
	}
	return "Accessories: " + strings.Join(strList, ", ")
}
