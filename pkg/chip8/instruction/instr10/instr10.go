// A package for adding a byte to a register.
package instr10

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// 7xkk - ADD Vx, byte
// Set Vx = Vx + kk.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	sum := state.Registers.V[args.X] + uint16(args.KK)
	state.Registers.V[args.X] = sum & 0x00FF
	register.NextInstruction(state.Registers)
}
