package stack

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// SP will always point to the next empty space on the stack
// which means it varies from 0x00 to 0x10
type Stack struct {
	data []uint16
}

func NewStack() *Stack {
	return &Stack{
		data: make([]uint16, 0, 0x10),
	}
}

func Push(stack *Stack, registers *register.RegisterFile, address uint16) {
	stack.data = append(stack.data, address)
	registers.SP++
}

func Pop(stack *Stack, registers *register.RegisterFile) uint16 {
	var address uint16

	registers.SP--
	address, stack.data = stack.data[len(stack.data)-1],
		stack.data[:len(stack.data)-1]

	return address
}
