// A package for making register receive another register.
package instr11

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
)

// 8xy0 - LD Vx, Vy
// Set Vx = Vy.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.V[args.X] = state.Registers.V[args.Y]
}
