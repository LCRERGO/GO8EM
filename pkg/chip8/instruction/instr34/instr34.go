package instr34

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/memory"
)

// Fx55 - LD [I], Vx
// Store registers V0 through Vx in memory starting at location I.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	for i := 0; i <= int(args.X); i++ {
		memory.Set(state.Memory,
			int(state.Registers.I)+i,
			uint8(state.Registers.V[i]))
	}
}
