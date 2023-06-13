package instr00

import (
	"log"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
)

// Tried to call an invalid instruction
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	log.Fatal("exec_instr00: invalid instruction")
}
