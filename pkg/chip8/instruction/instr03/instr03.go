package instr03

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
	"github.com/LCRERGO/GO8EM/pkg/chip8/stack"
)

// 00EE - RET
// Return from a subroutine.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.PC = stack.Pop(state.Stack, state.Registers)
}
