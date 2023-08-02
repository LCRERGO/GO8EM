// A package to load sprite given by Vx into I.
package instr32

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
	"github.com/LCRERGO/GO8EM/pkg/constants"
)

// Fx29 - LD F, Vx
// Set I = location of sprite for digit Vx.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.I = state.Registers.V[args.X] *
		constants.DefaultSpriteHeight
	register.NextInstruction(state.Registers)
}
