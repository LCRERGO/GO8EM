package instr31

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// Fx1E - ADD I, Vx
// Set I = I + Vx.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.I += state.Registers.V[args.X]
}
