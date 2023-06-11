package instr30

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// Fx18 - LD ST, Vx
// Set sound timer = Vx.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.ST = uint8(state.Registers.V[args.X])
}
