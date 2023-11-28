// A package for an error instruction.
package instr00

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/utils/log"
)

// Tried to call an invalid instruction
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	log.Fatal("exec_instr00: invalid instruction")
}
