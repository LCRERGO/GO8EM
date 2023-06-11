package instr07

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// 4xkk - SNE Vx, byte
// Skip next instruction if Vx != kk.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	if state.Registers.V[args.X] != uint16(args.KK) {
		register.NextInstruction(state.Registers)
	}
}
