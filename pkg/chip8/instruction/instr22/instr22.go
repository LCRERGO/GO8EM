package instr22

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// Bnnn - JP V0, addr
// Jump to location nnn + V0.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.PC = args.NNN + state.Registers.V[0x00]
}
