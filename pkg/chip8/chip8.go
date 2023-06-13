// Package of chip8 general purpose interface.
package chip8

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8/handler"
	"github.com/LCRERGO/GO8EM/pkg/chip8/keyboard"
	"github.com/LCRERGO/GO8EM/pkg/chip8/memory"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
	"github.com/LCRERGO/GO8EM/pkg/chip8/screen"
	"github.com/LCRERGO/GO8EM/pkg/chip8/stack"
)

// Chip8 entity is the logical representation of
// the components of the chip8
type Chip8 struct {
	Memory    *memory.Memory
	Stack     *stack.Stack
	Registers *register.RegisterFile
	Keyboard  *keyboard.Keyboard
	Screen    *screen.Screen

	Handler *handler.Handler
}

// Create a new Chip8
func New() *Chip8 {
	return &Chip8{
		Memory:    memory.New(),
		Stack:     stack.New(),
		Registers: register.New(),
		Keyboard:  keyboard.New(),
		Screen:    screen.New(),
	}
}

// Deep Copy a Chip8.
func Copy(chip8 *Chip8) *Chip8 {
	return &Chip8{
		Memory:    memory.Copy(chip8.Memory),
		Stack:     stack.Copy(chip8.Stack),
		Registers: register.Copy(chip8.Registers),
		Keyboard:  keyboard.Copy(chip8.Keyboard),
		Screen:    screen.Copy(chip8.Screen),

		Handler: handler.Copy(chip8.Handler),
	}
}

// Destroy a Chip8.
func Destroy(chip8 *Chip8) {
	screen.Destroy(chip8.Screen)
	keyboard.Destroy(chip8.Keyboard)
	register.Destroy(chip8.Registers)
	stack.Destroy(chip8.Stack)
	memory.Destroy(chip8.Memory)

	handler.Destroy(chip8.Handler)

	chip8 = nil
}

// Fetch an Opcode from Chip8.
func FetchOpcode(chip8 *Chip8) uint16 {
	var fstByte uint16 = uint16(memory.Get(chip8.Memory, int(chip8.Registers.PC)))
	var sndByte uint16 = uint16(memory.Get(chip8.Memory, int(chip8.Registers.PC)+1))

	return fstByte<<8 | sndByte
}

// Attach a Handler to Chip8.
func AddHandler(chip8 *Chip8, handler *handler.Handler) {
	chip8.Handler = handler
}
