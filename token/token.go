package token

// These are the defined tokens for the vm language
// That will need to be tranlated based on what the
// actual values of these segments are going to be

// Token is a type defined for what a parser will
// tranlate a string read from the file to.
type Token int

// Enum for tokens for VM code
const (
	PUSH Token = iota
	POP
	SP
	LCL
	ARG
	THIS
	THAT
	CONSTANT
	STATIC
	TEMP
	POINTER
)
