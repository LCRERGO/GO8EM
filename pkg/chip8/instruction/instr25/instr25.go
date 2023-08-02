// A package to skip a instruction given a keypress.
package instr25

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/keyboard"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// Ex9E - SKP Vx
// Skip next instruction if key with the value of Vx is pressed.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	key := state.Registers.V[args.X]
	if keyboard.IsDown(state.Keyboard, int(key)) {
		register.NextInstruction(state.Registers)
	}
	register.NextInstruction(state.Registers)
}
