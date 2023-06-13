// Package stack provides functions for stack manipulation.
package stack

import (
	"fmt"
	"log"

	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
)

// Stack entity that holds data for function return values
// SP will always point to the next empty space on the stack
// which means it varies from 0x00 to 0x10.
type Stack struct {
	data []uint16
}

// Create a new stack.
func New() *Stack {
	return &Stack{
		data: make([]uint16, 0x10),
	}
}

// Deep Copy a Stack.
func Copy(stack *Stack) *Stack {
	data := make([]uint16, 0x10)
	copy(data, stack.data)

	return &Stack{
		data: data,
	}
}

// Destroy a Stack.
func Destroy(stack *Stack) {
	stack.data = nil
	stack = nil
}

// Push return address of a function onto a Stack.
// A RegisterFile should be passed for altering SP properlly.
func Push(stack *Stack, registers *register.RegisterFile, address uint16) {
	if !isValidSP(registers.SP) {
		log.Fatal("stack_push: invalid sp")
	}
	stack.data[registers.SP] = address
	registers.SP++
}

// Pop return address of a function out of a Stack.
// A RegisterFile should be passed for altering SP properlly.
// Returns the value of return address of a function.
func Pop(stack *Stack, registers *register.RegisterFile) uint16 {
	var address uint16

	if !isValidSP(registers.SP) {
		log.Fatal("stack_pop: invalid sp")
	}
	address = stack.data[registers.SP-1]
	stack.data[registers.SP-1] = 0x00
	registers.SP--

	return address
}

// ToString returns the string representation of a Stack.
func ToString(stack *Stack) string {
	var str string

	for i, v := range stack.data {
		str += fmt.Sprintf("0x%02X: 0x%04X\n", i, v)
	}

	return str
}

func isValidSP(sp uint16) bool {
	return sp <= 0x10
}
