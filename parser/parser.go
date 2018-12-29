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
func New(f *os.File) *Parser {
	fiter := fileiterator.New(f)
	return &Parser{
		fiter: fiter,
	}
}

// HasMoreCommands returns true if there are more lines
// in the file that need to be read and interpreted
func (p *Parser) HasMoreCommands() bool {
	return p.fiter.HasNext()
}

// Parse the string given to the parser, return the tokens
// for interpretation to assembly, it might also return a
// nil slice
func (p *Parser) Parse(s string) []token.Token {
	s = removeComments(s)
	s = removeWhiteSpaces(s)
	return nil
}

// this function removes the comments from the string
func removeComments(s string) string {
	// A comment is the context of a vm file is going to
	// be a line which starts with //
	i := strings.Index(s, "//")
	if i != -1 {
		return s[:i]
	}
	return ""
}

// removeWhiteSpaces should remove all whitespace that is extraneous in the string,
// so that it can be passed down for further processing
func removeWhiteSpaces(s string) string {
	s = strings.TrimSpace(s)
	return removeExtraWhiteSpaces(s)
}

func removeExtraWhiteSpaces(s string) string {
	r := regexp.MustCompile(`\s+`)
	return r.ReplaceAllString(s, " ")
}
