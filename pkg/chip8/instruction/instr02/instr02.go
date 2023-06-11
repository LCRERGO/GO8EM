package instr02

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
	"github.com/LCRERGO/GO8EM/pkg/chip8/screen"
)

// 00E0 - CLS
// Clear the display.
func Exec(state *chip8.Chip8, args *instruction.OpcodeArguments) {
	screen.Clear(state.Screen)
}
