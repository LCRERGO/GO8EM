// A package for jump to address.
package instr01

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
)

// 0nnn - SYS addr
// Jump to a machine code routine at nnn.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.PC = args.NNN
}
