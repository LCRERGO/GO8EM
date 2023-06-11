package instr17

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// 8xy6 - SHR Vx {, Vy}
// Set Vx = Vx SHR 1.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	state.Registers.V[0x0F] = state.Registers.V[args.X] & 0x01
	state.Registers.V[args.X] >>= 1
}
