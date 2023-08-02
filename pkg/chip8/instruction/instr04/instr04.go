// A package for a jump to address.
package instr04

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
)

// 1nnn - JP addr
// Jump to location nnn.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.PC = args.NNN
}
