package instr29

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
)

// Fx15 - LD DT, Vx
// Set delay timer = Vx.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.DT = uint8(state.Registers.V[args.X])
}
