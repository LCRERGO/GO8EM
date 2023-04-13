package stack

import "github.com/LCRERGO/GO8EM/pkg/chip8/register"

type Stack struct {
	data [0xF]uint16
}

func Push(stack *Stack, registers *register.RegisterFile, address uint16) {
	registers.SP++
	stack.data[registers.SP] = address
}

func Pop(stack *Stack, registers *register.RegisterFile) uint16 {
	address := stack.data[registers.SP]
	registers.SP--

	return address
}
