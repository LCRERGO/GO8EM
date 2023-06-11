package instr08

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// 5xy0 - SE Vx, Vy
// Skip next instruction if Vx = Vy.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	if state.Registers.V[args.X] == state.Registers.V[args.Y] {
		register.NextInstruction(state.Registers)
	}
}