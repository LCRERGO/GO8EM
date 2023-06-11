package instr14

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// 8xy3 - XOR Vx, Vy
// Set Vx = Vx XOR Vy.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.V[args.X] ^= state.Registers.V[args.Y]
}