package octocat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccessorySets(t *testing.T) {
	hat := NewAccessory("hat")
	coat := NewAccessory("coat")
	flower := NewAccessory("flower")
	accessories := []*Accessory{hat, coat, flower}

	result := GetAccessorySets(accessories)

	assert.Equal(t, 8, len(result))
	assertIncludesSet(t, result, "no accessories")
	assertIncludesSet(t, result, "Accessories: hat")
	assertIncludesSet(t, result, "Accessories: coat")
	assertIncludesSet(t, result, "Accessories: flower")
	assertIncludesSet(t, result, "Accessories: hat, coat")
	assertIncludesSet(t, result, "Accessories: hat, coat, flower")
	assertIncludesSet(t, result, "Accessories: hat, flower")
	assertIncludesSet(t, result, "Accessories: coat, flower")
}

func assertIncludesSet(t *testing.T, haystack []*AccessorySet, needle string) {
	found := false
	for _, haystackEl := range haystack {
		if haystackEl.String() == needle {
			found = true
			break
		}
	}
	assert.True(t, found, "could not find: "+needle)
}
