// A package for adding two registers.
package instr15

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// 8xy4 - ADD Vx, Vy
// Set Vx = Vx + Vy, set VF = carry.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	result := state.Registers.V[args.X] + state.Registers.V[args.Y]

	if result > 0xFF {
		state.Registers.V[0x0F] = 0x01
		result &= 0xFF
	} else {
		state.Registers.V[0x0F] = 0x00
	}
	if args.X != 0x0F {
		state.Registers.V[args.X] = result
	}
	register.NextInstruction(state.Registers)
}
