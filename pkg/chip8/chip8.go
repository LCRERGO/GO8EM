package chip8

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8/memory"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
	"github.com/LCRERGO/GO8EM/pkg/chip8/screen"
	"github.com/LCRERGO/GO8EM/pkg/chip8/stack"
)

type Chip8 struct {
	Memory    *memory.Memory
	Stack     *stack.Stack
	Registers *register.RegisterFile
	Screen    *screen.Screen
}

func New() *Chip8 {
	return &Chip8{
		Memory:    memory.New(),
		Stack:     stack.New(),
		Registers: register.New(),
		Screen:    screen.New(),
	}
}

func Destroy(chip8 *Chip8) {
	screen.Destroy(chip8.Screen)
	register.Destroy(chip8.Registers)
	stack.Destroy(chip8.Stack)
	memory.Destroy(chip8.Memory)

	chip8 = nil
}
