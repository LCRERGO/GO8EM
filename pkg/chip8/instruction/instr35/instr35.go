// A package to load registers V0 to Vx at [I].
package instr35

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/memory"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// Fx65 - LD Vx, [I]
// Read registers V0 through Vx from memory starting at location I.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	for i := 0; i <= int(args.X); i++ {
		state.Registers.V[i] = uint16(memory.Get(state.Memory, int(state.Registers.I)+i))
	}
	register.NextInstruction(state.Registers)
}
