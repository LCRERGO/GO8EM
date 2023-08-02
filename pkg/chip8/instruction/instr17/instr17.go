// A package for shifting right a register.
package instr17

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// 8xy6 - SHR Vx {, Vy}
// Set Vx = Vx SHR 1.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.V[0x0F] = state.Registers.V[args.X] & 0x01
	if args.X != 0x0F {
		state.Registers.V[args.X] >>= 1
	}
	register.NextInstruction(state.Registers)
}
