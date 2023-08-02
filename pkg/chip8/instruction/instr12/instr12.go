// A package for "ORing" two registers.
package instr12

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// 8xy1 - OR Vx, Vy
// Set Vx = Vx OR Vy.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.V[args.X] |= state.Registers.V[args.Y]
	state.Registers.V[0x0F] = 0x00
	register.NextInstruction(state.Registers)
}
