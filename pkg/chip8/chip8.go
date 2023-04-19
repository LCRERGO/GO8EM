package chip8

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8/memory"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
	"github.com/LCRERGO/GO8EM/pkg/chip8/screen"
	"github.com/LCRERGO/GO8EM/pkg/chip8/stack"
)

type Chip8 struct {
	Memory    memory.Memory
	Stack     *stack.Stack
	Registers register.RegisterFile
	Screen    *screen.Screen
}

func NewChip8() Chip8 {
	return Chip8{
		Stack:  stack.NewStack(),
		Screen: screen.NewScreen(),
	}
}
