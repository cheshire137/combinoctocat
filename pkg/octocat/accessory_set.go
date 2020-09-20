package octocat

import (
	"strings"

	set "github.com/deckarep/golang-set"
)

type AccessorySet struct {
	Accessories []*Accessory
}

func NoAccessories() *AccessorySet {
	return &AccessorySet{}
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

func GetAccessorySets(accessoryChoices []*Accessory) []*AccessorySet {
	interfaceSlice := make([]interface{}, len(accessoryChoices))
	for i, accessoryChoice := range accessoryChoices {
		interfaceSlice[i] = accessoryChoice
	}
	allSubsets := set.NewSetFromSlice(interfaceSlice).PowerSet()
	accessorySets := []*AccessorySet{NoAccessories()}
	for el := range allSubsets.Iter() {
		subset, ok := el.(set.Set)
		if ok {
			elsInSet := subset.ToSlice()
			if len(elsInSet) < 1 {
				continue
			}

			accessoriesInSet := make([]*Accessory, len(elsInSet))
			for i, elInSet := range elsInSet {
				accessoryInSet, ok := elInSet.(*Accessory)
				if ok {
					accessoriesInSet[i] = accessoryInSet
				}
			}
			accessorySets = append(accessorySets, NewAccessorySet(accessoriesInSet...))
		}
	}
	return accessorySets
}
