package instr13

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// 8xy2 - AND Vx, Vy
// Set Vx = Vx AND Vy.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.V[args.X] &= state.Registers.V[args.Y]
}