package instr26

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
	"github.com/LCRERGO/GO8EM/pkg/chip8/keyboard"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// ExA1 - SKNP Vx
// Skip next instruction if key with the value of Vx is not pressed.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	key := state.Registers.V[args.X]
	if keyboard.IsUp(state.Keyboard, int(key)) {
		register.NextInstruction(state.Registers)
	}
}