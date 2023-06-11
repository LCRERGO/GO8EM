package instr05

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
	"github.com/LCRERGO/GO8EM/pkg/chip8/stack"
)

// 2nnn - CALL addr
// Call subroutine at nnn.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	stack.Push(state.Stack, state.Registers, state.Registers.PC)
	state.Registers.PC = args.NNN
}
