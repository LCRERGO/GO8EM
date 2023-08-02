// A package to load an address to I register.
package instr21

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// Annn - LD I, addr
// Set I = nnn.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.I = args.NNN
	register.NextInstruction(state.Registers)
}
