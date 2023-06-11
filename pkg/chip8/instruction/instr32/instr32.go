package instr32

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
	"github.com/LCRERGO/GO8EM/pkg/constants"
)

// Fx29 - LD F, Vx
// Set I = location of sprite for digit Vx.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.I = state.Registers.V[args.X] *
		constants.DefaultSpriteHeight
}
