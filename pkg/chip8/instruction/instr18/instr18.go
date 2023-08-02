// A package for negative subtracting two registers.
package instr18

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
)

// 8xy7 - SUBN Vx, Vy
// Set Vx = Vy - Vx, set VF = NOT borrow.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	var result uint8 = uint8(state.Registers.V[args.Y] - state.Registers.V[args.X])

	if state.Registers.V[args.Y] < state.Registers.V[args.X] {
		state.Registers.V[0x0F] = 0x01
	} else {
		state.Registers.V[0x0F] = 0x00
	}
	state.Registers.V[args.X] = uint16(result)
}
