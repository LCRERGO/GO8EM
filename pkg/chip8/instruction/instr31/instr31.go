// A package to add Vx into I.
package instr31

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// Fx1E - ADD I, Vx
// Set I = I + Vx.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.I = (state.Registers.I + state.Registers.V[args.X]) & 0xFFF
	register.NextInstruction(state.Registers)
}
