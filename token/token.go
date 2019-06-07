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
	LCL Segment = iota
	ARG
	THIS
	THAT
	CONSTANT
	STATIC
	TEMP
	POINTER
)

// TokenMap holds all possible keyword to token mapping
var TokenMap = map[string]Token{
	"add":  ARITHMETIC,
	"sub":  ARITHMETIC,
	"neg":  ARITHMETIC,
	"eq":   ARITHMETIC,
	"gt":   ARITHMETIC,
	"lt":   ARITHMETIC,
	"and":  ARITHMETIC,
	"or":   ARITHMETIC,
	"not":  ARITHMETIC,
	"push": PUSH,
	"pop":  POP,
}

// SegmentMap holds all possible keyword to segment mapping
var SegmentMap = map[string]Segment{
	"local":    LCL,
	"argument": ARG,
	"static":   STATIC,
	"constant": CONSTANT,
	"this":     THIS,
	"that":     THAT,
	"temp":     TEMP,
	"pointer":  POINTER,
}

// Command covers all vm commands
type Command struct {
	T    Token
	Arg1 string
	Arg2 string
	S    string
}
