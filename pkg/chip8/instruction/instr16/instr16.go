// A package for subtracting two registers.
package instr16

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// 8xy5 - SUB Vx, Vy
// Set Vx = Vx - Vy, set VF = NOT borrow.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	var result uint8 = uint8(state.Registers.V[args.X] - state.Registers.V[args.Y])

	if state.Registers.V[args.X] > state.Registers.V[args.Y] {
		state.Registers.V[0x0F] = 0x01
	} else {
		state.Registers.V[0x0F] = 0x00
	}
	if args.X != 0x0F {
		state.Registers.V[args.X] = uint16(result)
	}
	register.NextInstruction(state.Registers)
}
