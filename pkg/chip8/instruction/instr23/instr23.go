package instr23

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/utils/lcg"
)

// Cxkk - RND Vx, byte
// Set Vx = random byte AND kk.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.V[args.X] =
		uint16((lcg.LCG() & int(args.KK)) & 0xFF)
}
