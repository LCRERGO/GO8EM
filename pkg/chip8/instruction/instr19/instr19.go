package instr19

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// 8xyE - SHL Vx {, Vy}
// Set Vx = Vx SHL 1.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.V[0x0F] = state.Registers.V[args.X] & (0x01 << 7)
	state.Registers.V[args.X] = (state.Registers.V[args.X] << 1) & 0xFF
}