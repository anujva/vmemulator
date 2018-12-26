package parser

import "go/token"

// Parser for vm code
type Parser struct {
}

// Parse the string given to the parser, return the tokens
// for interpretation to assembly
func (p *Parser) Parse(s string) []token.Token {

}

// this function removes the comments from the string
func removeComments(s string) string {

}
