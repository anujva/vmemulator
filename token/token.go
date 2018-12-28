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
	ARITHMETIC
)

// Segment defines the memory segment which is the
// entry following PUSH/POP command
type Segment int

// Enum for memory segments
const (
	SP Segment = iota
	LCL
	ARG
	THIS
	THAT
	CONSTANT
	STATIC
	TEMP
	POINTER
)
