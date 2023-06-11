package instr01

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// 0nnn - SYS addr
// Jump to a machine code routine at nnn.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.PC = args.NNN
}
