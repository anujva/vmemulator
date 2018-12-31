package codewriter

import (
	"fmt"

	"github.com/anujva/vmemulator/token"
)

// CodeWriter will be used to generate the assembly
type CodeWriter struct {
}

// GenerateAssembly returns the string representation of the assembly code
func (cw *CodeWriter) GenerateAssembly(command token.Command) string {
	switch command.T {
	case token.ARITHMETIC:
		return fmt.Sprintf(arithmeticTemplate)
	}
}

var arithmeticTemplate = `
%s
@%s
D=A
@SP
A=M
M=D
@SP
M=M+1
`
