package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPositiveEnglishNumberName(t *testing.T) {
	assert.Equal(t, "ten", GetPositiveEnglishNumberName(10))
	assert.Equal(t, "fifty-one", GetPositiveEnglishNumberName(51))
	assert.Equal(t, "eight hundred thirty-nine", GetPositiveEnglishNumberName(839))
	assert.Equal(t, "two thousand", GetPositiveEnglishNumberName(2000))
	assert.Equal(t, "one thousand two hundred thirty-four", GetPositiveEnglishNumberName(1234))
	assert.Equal(t, "twelve thousand three hundred forty-five", GetPositiveEnglishNumberName(12345))
	assert.Equal(t, "one hundred twenty-three thousand four hundred fifty-six", GetPositiveEnglishNumberName(123456))
	assert.Equal(t, "one million two hundred thirty-four thousand five hundred sixty-seven", GetPositiveEnglishNumberName(1234567))
	assert.Equal(t, "one billion two hundred thirty-four million five hundred sixty-seven thousand eight hundred ninety", GetPositiveEnglishNumberName(1234567890))
	assert.Equal(t, "one trillion two hundred thirty-four billion five hundred sixty-seven million eight hundred ninety thousand one hundred twenty-three", GetPositiveEnglishNumberName(1234567890123))
	assert.Equal(t, "one quadrillion two hundred thirty-four trillion five hundred sixty-seven billion eight hundred ninety million one hundred twenty-three thousand four hundred fifty-six", GetPositiveEnglishNumberName(1234567890123456))
	assert.Equal(t, "one quintrillion two hundred thirty-four quadrillion five hundred sixty-seven trillion eight hundred ninety billion one hundred twenty-three million four hundred fifty-six thousand seven hundred eighty-nine", GetPositiveEnglishNumberName(1234567890123456789))
	assert.Equal(t, "three quintrillion eight hundred sixty-nine quadrillion eight hundred three trillion seven hundred seven billion three hundred sixty-five million three hundred twenty-six thousand eight hundred forty-eight", GetPositiveEnglishNumberName(3869803707365326848))
}
