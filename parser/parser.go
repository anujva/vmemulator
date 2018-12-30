package parser

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/anujva/iterator"
	"github.com/anujva/iterator/fileiterator"
	"github.com/anujva/vmemulator/token"
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
func (p *Parser) Parse(ctx context.Context) (*token.Command, error) {
	s := p.fiter.Next(ctx).(*string)
	fmt.Println("Debugging: ", *s)
	str := removeComments(*s)
	str = removeWhiteSpaces(str)
	// split the command, and generate the tokens
	if str == "" {
		return nil, nil
	}
	sp := strings.Split(str, " ")
	return convertToTokens(sp)
}

func convertToTokens(sp []string) (*token.Command, error) {
	// this will try to convert the string to command
	var val token.Token
	var ok bool
	fmt.Println("Debugging: ", sp[0])
	if val, ok = token.TokenMap[sp[0]]; !ok {
		return nil, errors.New("unidentified command")
	}

	switch val {
	case token.ARITHMETIC:
		return &token.Command{
			T:    val,
			Arg1: sp[0],
		}, nil
	case token.PUSH:
		return &token.Command{
			T:    val,
			Arg1: sp[1],
			Arg2: sp[2],
		}, nil
	case token.POP:
		return &token.Command{
			T:    val,
			Arg1: sp[1],
			Arg2: sp[2],
		}, nil
	default:
		return nil, errors.New("could not decode token")
	}
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
