// A package to draw a sprite to screen.
package instr24

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/memory"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
	"github.com/LCRERGO/GO8EM/pkg/chip8/screen"
)

// Dxyn - DRW Vx, Vy, nibble
// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	sprite := memory.FetchSprite(state.Memory, int(state.Registers.I), int(args.N))
	collision := screen.DrawSprite(state.Screen,
		int(state.Registers.V[args.X]), int(state.Registers.V[args.Y]),
		sprite,
		int(args.N))
	if collision {
		state.Registers.V[0x0F] = 0x01
	} else {
		state.Registers.V[0x0F] = 0x00
	}
	register.NextInstruction(state.Registers)
}
