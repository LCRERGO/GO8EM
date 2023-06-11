package instr27

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// Fx07 - LD Vx, DT
// Set Vx = delay timer value.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.V[args.X] = uint16(state.Registers.DT)
}