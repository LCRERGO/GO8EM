// A package to clear the screen.
package instr02

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/screen"
)

// 00E0 - CLS
// Clear the display.
func Exec(state *chip8.Chip8, args *argument.OpcodeArguments) {
	screen.Clear(state.Screen)
}
