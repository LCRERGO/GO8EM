package instr16

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
)

// 8xy5 - SUB Vx, Vy
// Set Vx = Vx - Vy, set VF = NOT borrow.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	result := state.Registers.V[args.X] - state.Registers.V[args.Y]

	if state.Registers.V[args.X] > state.Registers.V[args.Y] {
		state.Registers.V[0x0F] = 0x01
	} else {
		state.Registers.V[0x0F] = 0x00
	}
	state.Registers.V[args.X] = result
}
