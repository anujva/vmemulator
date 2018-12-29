package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveComments(t *testing.T) {
	s := "this is a command // followed by comment"
	scommentremoved := removeComments(s)
	assert.Equal(t, "this is a command ", scommentremoved, "The two strings should be same")
}

func TestRemoveWhitespaces(t *testing.T) {
	s := "this is another hard command, "
	swhitespaceremoved := removeWhiteSpaces(s)
	assert.Equal(t, "this is another hard command,", swhitespaceremoved,
		"The two strings should be equal")

	s1 := "       this is    another string    with extra spaces   "
	sremoved := removeWhiteSpaces(s1)
	assert.Equal(t, "this is another string with extra spaces",
		sremoved, "The two strings should be equal")
}
