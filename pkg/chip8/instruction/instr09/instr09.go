// A package to set a register as a byte.
package instr09

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// 6xkk - LD Vx, byte
// Set Vx = kk.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.V[args.X] = uint16(args.KK)
	register.NextInstruction(state.Registers)
}
