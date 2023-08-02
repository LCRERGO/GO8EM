// A package to load Vx into sound timer.
package instr30

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// Fx18 - LD ST, Vx
// Set sound timer = Vx.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	state.Registers.ST = uint8(state.Registers.V[args.X])
	register.NextInstruction(state.Registers)
}
