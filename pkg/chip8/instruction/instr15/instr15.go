package instr15

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
)

// 8xy4 - ADD Vx, Vy
// Set Vx = Vx + Vy, set VF = carry.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	result := state.Registers.V[args.X] + state.Registers.V[args.Y]

	if result > 0xFF {
		state.Registers.V[0x0F] = 0x01
		result &= 0xFF
	} else {
		state.Registers.V[0x0F] = 0x00
	}
	state.Registers.V[args.X] = result
}
