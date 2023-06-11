package instr21

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// Annn - LD I, addr
// Set I = nnn.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.I = args.NNN
}
