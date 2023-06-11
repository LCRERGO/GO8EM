package instr10

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// 7xkk - ADD Vx, byte
// Set Vx = Vx + kk.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.V[args.X] += uint16(args.KK)
}