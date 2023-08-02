// A package to load the BCD representation starting at Vx.
package instr33

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/memory"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
	"github.com/LCRERGO/GO8EM/pkg/utils"
)

// Fx33 - LD B, Vx
// Store BCD representation of Vx in memory locations I, I+1, and I+2.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	units, tens, hundreds := utils.Byte2BCD(uint8(state.Registers.V[args.X]))

	memory.Set(state.Memory, int(state.Registers.I), hundreds)
	memory.Set(state.Memory, int(state.Registers.I+1), tens)
	memory.Set(state.Memory, int(state.Registers.I+2), units)
	register.NextInstruction(state.Registers)
}
