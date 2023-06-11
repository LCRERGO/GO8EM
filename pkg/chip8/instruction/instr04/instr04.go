package instr04

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// 1nnn - JP addr
// Jump to location nnn.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.PC = args.NNN
}
