package octocat

import (
	"fmt"
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
	fmt.Println(result)
	assertIncludesSet(t, result, "no accessories")
	assertIncludesSet(t, result, "Accessories (1): hat")
	assertIncludesSet(t, result, "Accessories (1): coat")
	assertIncludesSet(t, result, "Accessories (1): flower")
	assertIncludesSet(t, result, "Accessories (2): coat, hat")
	assertIncludesSet(t, result, "Accessories (3): coat, flower, hat")
	assertIncludesSet(t, result, "Accessories (2): flower, hat")
	assertIncludesSet(t, result, "Accessories (2): coat, flower")
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
