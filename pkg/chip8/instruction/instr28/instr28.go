// A package to load a key into Vx.
package instr28

import (
	"log"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// Fx0A - LD Vx, K
// Wait for a key press, store the value of the key in Vx.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	key, err := state.Handler.WaitForKeyPress()
	if err != nil {
		log.Fatal("exec_instr28: wait_for_key_press")
	}
	state.Registers.V[args.X] = uint16(key)
	register.NextInstruction(state.Registers)
}
