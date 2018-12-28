package parser

import (
	"go/token"
	"os"
	"regexp"
	"strings"

	"github.com/anujva/iterator"
	"github.com/anujva/iterator/fileiterator"
)

// Parser for vm code
type Parser struct {
	fiter iterator.Iterator
}

// New returns a pointer to Parser
func New(f os.File) *Parser {
	fiter := fileiterator.New(f)
	return &Parser{
		fiter: fiter,
	}
}

// Parse the string given to the parser, return the tokens
// for interpretation to assembly
func (p *Parser) Parse(s string) []token.Token {

}

// this function removes the comments from the string
func removeComments(s string) string {
	// A comment is the context of a vm file is going to
	// be a line which starts with //

}

// removeWhiteSpaces should remove all whitespace that is extraneous in the string,
// so that it can be passed down for further processing
func removeWhiteSpaces(s string) string {
	s := strings.TrimSpace(s)
	return removeExtraWhiteSpaces(s)
}

func removeExtraWhiteSpaces(s string) string {
	r := regexp.MustCompile(`\s+`)
	return r.ReplaceAllString(s, " ")
}
